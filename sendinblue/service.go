package sendinblue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/visiperf/visimail"
)

type service struct {
	apiKey     string
	apiVersion string
	baseURL    string
}

// NewService is func to create new instance of sendinblue service implementing email service interface
func NewService(apiKey, apiVersion, baseURL string) visimail.EmailService {
	return &service{
		apiKey:     apiKey,
		apiVersion: apiVersion,
		baseURL:    baseURL,
	}
}

func (s *service) SendEmailFromTemplate(templateID string, to []string, cc []string, bcc []string, params map[string]interface{}, attachments []*visimail.Attachment, replyTo string) (string, error) {
	t, err := strconv.Atoi(templateID)
	if err != nil {
		return "", &ParamsError{fmt.Errorf("failed to convert template id to int: %v", err)}
	}

	var as []*attachment
	for _, attachment := range attachments {
		a, err := newAttachment(attachment.URL, attachment.Content, attachment.Name)
		if err != nil {
			return "", &ParamsError{fmt.Errorf("failed to create new attachment: %v", err)}
		}

		as = append(as, a)
	}

	var payload = struct {
		TemplateID  int64                  `json:"templateId"`
		To          []*destination         `json:"to"`
		Cc          []*destination         `json:"cc,omitempty"`
		Bcc         []*destination         `json:"bcc,omitempty"`
		Params      map[string]interface{} `json:"params,omitempty"`
		Attachments []*attachment          `json:"attachment,omitempty"`
		ReplyTo     *destination           `json:"replyTo,omitempty"`
	}{
		TemplateID:  int64(t),
		To:          newDestinations(to),
		Cc:          newDestinations(cc),
		Bcc:         newDestinations(bcc),
		Params:      params,
		Attachments: as,
		ReplyTo:     newDestination(replyTo),
	}

	return s.post(payload, "/smtp/email")
}

func (s *service) post(payload interface{}, endpoint string) (string, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s%s", s.baseURL, s.apiVersion, endpoint), bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", s.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var e QueryError
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return "", err
		}

		return "", &e
	}

	var r struct {
		MessageID string `json:"messageId"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", err
	}

	return r.MessageID, nil
}
