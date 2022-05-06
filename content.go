package visimail

import (
	"encoding/json"
	"errors"
)

var (
	ErrEmptyHTMLContent      = errors.New("html content is empty")
	ErrEmptyPlainTextContent = errors.New("plain text content is empty")
	ErrEmptyTemplateId       = errors.New("template id is empty")
)

type content interface {
	json.Marshaler
	Validate() error
}

var emptyHTMLContent = htmlContent{}

type htmlContent struct {
	html string
}

func newHTMLContent(html string) htmlContent {
	return htmlContent{html}
}

func (c htmlContent) Content() string {
	return c.html
}

func (c htmlContent) Validate() error {
	if c.IsZero() {
		return ErrEmptyHTMLContent
	}

	return nil
}

func (c htmlContent) Equals(content htmlContent) bool {
	return c == content
}

func (c htmlContent) IsZero() bool {
	return c.Equals(emptyHTMLContent)
}

func (c htmlContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Content string `json:"htmlContent"`
	}{
		Content: c.Content(),
	})
}

var emptyPlainTextContent = plainTextContent{}

type plainTextContent struct {
	text string
}

func newPlainTextContent(text string) plainTextContent {
	return plainTextContent{text}
}

func (c plainTextContent) Content() string {
	return c.text
}

func (c plainTextContent) Validate() error {
	if c.IsZero() {
		return ErrEmptyPlainTextContent
	}

	return nil
}

func (c plainTextContent) Equals(content plainTextContent) bool {
	return c == content
}

func (c plainTextContent) IsZero() bool {
	return c.Equals(emptyPlainTextContent)
}

func (c plainTextContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Content string `json:"textContent"`
	}{
		Content: c.Content(),
	})
}
