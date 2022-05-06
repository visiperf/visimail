package visimail

import (
	"encoding/json"
	"errors"
	"reflect"
)

var (
	ErrEmptyHTMLContent      = errors.New("html content is empty")
	ErrEmptyPlainTextContent = errors.New("plain text content is empty")
	ErrEmptyTemplateId       = errors.New("template id is empty")
)

type contentType int

const (
	contentTypeHTML contentType = iota
	contentTypePlainText
	contentTypeTemplateId
)

type content interface {
	json.Marshaler
	Validate() error
	Type() contentType
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

func (c htmlContent) Type() contentType {
	return contentTypeHTML
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

func (c plainTextContent) Type() contentType {
	return contentTypePlainText
}

var emptyTemplateIdContent = templateIdContent{}

type templateIdContent struct {
	id     int
	params map[string]interface{}
}

func newTemplateIdContent(templateId int, opts ...TemplateIdContentOption) templateIdContent {
	c := templateIdContent{id: templateId}
	for _, opt := range opts {
		opt(&c)
	}

	return c
}

func (c templateIdContent) ID() int {
	return c.id
}

func (c templateIdContent) Params() map[string]interface{} {
	return c.params
}

func (c templateIdContent) Validate() error {
	if c.IsZero() {
		return ErrEmptyTemplateId
	}

	return nil
}

func (c templateIdContent) Equals(content templateIdContent) bool {
	return reflect.DeepEqual(c, content)
}

func (c templateIdContent) IsZero() bool {
	return c.Equals(emptyTemplateIdContent)
}

func (c templateIdContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TemplateId int                    `json:"templateId"`
		Params     map[string]interface{} `json:"params,omitempty"`
	}{
		TemplateId: c.ID(),
		Params:     c.Params(),
	})
}

func (c templateIdContent) Type() contentType {
	return contentTypeTemplateId
}

type TemplateIdContentOption func(c *templateIdContent)

func WithParams(params map[string]interface{}) TemplateIdContentOption {
	return func(c *templateIdContent) {
		c.params = params
	}
}
