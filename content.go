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

type ContentType int

const (
	ContentTypeHTML ContentType = iota
	ContentTypePlainText
	ContentTypeTemplate
)

type Content interface {
	json.Marshaler
	Type() ContentType
	Validate() error
}

var emptyHTMLContent = HTMLContent{}

type HTMLContent struct {
	html string
}

func NewHTMLContent(html string) HTMLContent {
	return HTMLContent{html}
}

func (c HTMLContent) Content() string {
	return c.html
}

func (c HTMLContent) Equals(content HTMLContent) bool {
	return c == content
}

func (c HTMLContent) IsZero() bool {
	return c.Equals(emptyHTMLContent)
}

func (c HTMLContent) Type() ContentType {
	return ContentTypeHTML
}

func (c HTMLContent) Validate() error {
	if c.IsZero() {
		return ErrEmptyHTMLContent
	}

	return nil
}

func (c HTMLContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Content string `json:"htmlContent"`
	}{
		Content: c.Content(),
	})
}

var emptyPlainTextContent = PlainTextContent{}

type PlainTextContent struct {
	text string
}

func NewPlainTextContent(text string) PlainTextContent {
	return PlainTextContent{text}
}

func (c PlainTextContent) Content() string {
	return c.text
}

func (c PlainTextContent) Equals(content PlainTextContent) bool {
	return c == content
}

func (c PlainTextContent) IsZero() bool {
	return c.Equals(emptyPlainTextContent)
}

func (c PlainTextContent) Type() ContentType {
	return ContentTypePlainText
}

func (c PlainTextContent) Validate() error {
	if c.IsZero() {
		return ErrEmptyPlainTextContent
	}

	return nil
}

func (c PlainTextContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Content string `json:"textContent"`
	}{
		Content: c.Content(),
	})
}

type TemplateContent struct {
	id     int
	params map[string]interface{}
}

func NewTemplateContent(id int, opts ...TemplateContentOption) TemplateContent {
	c := TemplateContent{id: id}
	for _, opt := range opts {
		opt(&c)
	}

	return c
}

func (c TemplateContent) ID() int {
	return c.id
}

func (c TemplateContent) Params() map[string]interface{} {
	return c.params
}

func (c TemplateContent) Type() ContentType {
	return ContentTypeTemplate
}

func (c TemplateContent) Validate() error {
	if c.ID() <= 0 {
		return ErrEmptyTemplateId
	}

	return nil
}

func (c TemplateContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TemplateId int                    `json:"templateId"`
		Params     map[string]interface{} `json:"params,omitempty"`
	}{
		TemplateId: c.ID(),
		Params:     c.Params(),
	})
}

type TemplateContentOption func(c *TemplateContent)

func WithParams(params map[string]interface{}) TemplateContentOption {
	return func(c *TemplateContent) {
		c.params = params
	}
}
