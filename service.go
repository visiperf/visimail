package visimail

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL    = "https://api.sendinblue.com"
	apiVersion = "v3"
)

type Service struct {
	defaultHeaders map[string]string
}

func NewService() *Service {
	return &Service{
		defaultHeaders: map[string]string{
			"api-key":      env.Sendinblue.ApiKey,
			"Accept":       "application/json",
			"Content-Type": "application/json",
		},
	}
}

// TODO: add Service.SendChunkedEmail(ctx context.Context, email *Email, nbAttachmentsPerChunk int) (chan string, chan error)

func (s *Service) SendEmail(ctx context.Context, email *Email) (string, error) {
	return s.sendEmail(ctx, email)
}

func (s *Service) sendEmail(_ context.Context, email *Email) (string, error) {
	if err := email.Validate(); err != nil {
		return "", err
	}

	payload, err := json.Marshal(email)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, s.requestURL("/smtp/email"), bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	s.applyHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if isHttpError(resp.StatusCode) {
		var he HttpError
		if err := json.Unmarshal(body, &he); err != nil {
			return "", err
		}

		return "", he
	}

	var obj struct {
		MessageID string `json:"messageId"`
	}
	if err := json.Unmarshal(body, &obj); err != nil {
		return "", err
	}

	return obj.MessageID, nil
}

func (s *Service) requestURL(endpoint string) string {
	return fmt.Sprintf("%s/%s%s", baseURL, apiVersion, endpoint)
}

func (s *Service) applyHeaders(request *http.Request) {
	for key, value := range s.defaultHeaders {
		request.Header.Set(key, value)
	}
}
