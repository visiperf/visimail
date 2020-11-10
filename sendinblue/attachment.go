package sendinblue

import (
	"errors"
	"fmt"
)

type attachment struct {
	URL     string `json:"url,omitempty"`
	Content string `json:"content,omitempty"`
	Name    string `json:"name,omitempty"`
}

func newAttachment(url, content, name string) (*attachment, error) {
	attachment := &attachment{URL: url, Content: content, Name: name}
	if err := attachment.validate(); err != nil {
		return nil, fmt.Errorf("invalid attachment: %w", err)
	}

	return attachment, nil
}

func (a *attachment) validate() error {
	if a.URL != "" {
		return nil
	}

	if a.Content == "" {
		return errors.New("url or content is required")
	}

	if a.Name == "" {
		return errors.New("name is required when content is set")
	}

	return nil
}
