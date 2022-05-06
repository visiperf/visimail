package visimail

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
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
	ContentTypeTemplateId
)

type Content interface {
	Value() string
	Type() ContentType
	Validate() error
	IsZero() bool
}

var emptyHTMLContent = HTMLContent{}

type HTMLContent struct {
	html string
}

func NewHTMLContent(html string) HTMLContent {
	return HTMLContent{html}
}

func (c HTMLContent) HTML() string {
	return c.html
}

func (c HTMLContent) Value() string {
	return c.HTML()
}

func (c HTMLContent) Type() ContentType {
	return ContentTypeHTML
}

func (c HTMLContent) Validate() error {
	if len(c.html) <= 0 {
		return ErrEmptyHTMLContent
	}

	return nil
}

func (c HTMLContent) Equals(content HTMLContent) bool {
	return c == content
}

func (c HTMLContent) IsZero() bool {
	return c.Equals(emptyHTMLContent)
}

func (c HTMLContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Content string `json:"htmlContent"`
	}{
		Content: c.Value(),
	})
}

var emptyPlainTextContent = PlainTextContent{}

type PlainTextContent struct {
	text string
}

func NewPlainTextContent(text string) PlainTextContent {
	return PlainTextContent{text}
}

func (c PlainTextContent) Text() string {
	return c.text
}

func (c PlainTextContent) Value() string {
	return c.Text()
}

func (c PlainTextContent) Type() ContentType {
	return ContentTypePlainText
}

func (c PlainTextContent) Validate() error {
	if len(c.text) <= 0 {
		return ErrEmptyPlainTextContent
	}

	return nil
}

func (c PlainTextContent) Equals(content PlainTextContent) bool {
	return c == content
}

func (c PlainTextContent) IsZero() bool {
	return c.Equals(emptyPlainTextContent)
}

func (c PlainTextContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Content string `json:"textContent"`
	}{
		Content: c.Value(),
	})
}

var emptyTemplateIdContent = TemplateIdContent{}

type TemplateIdContent struct {
	id     int
	params map[string]interface{}
}

func NewTemplateIdContent(templateId int, opts ...TemplateIdContentOption) TemplateIdContent {
	c := TemplateIdContent{id: templateId}
	for _, opt := range opts {
		opt(&c)
	}

	return c
}

func (c TemplateIdContent) TemplateId() int {
	return c.id
}

func (c TemplateIdContent) Params() map[string]interface{} {
	return c.params
}

func (c TemplateIdContent) Value() string {
	return strconv.Itoa(c.TemplateId())
}

func (c TemplateIdContent) Type() ContentType {
	return ContentTypeTemplateId
}

func (c TemplateIdContent) Validate() error {
	if c.id <= 0 {
		return ErrEmptyTemplateId
	}

	return nil
}

func (c TemplateIdContent) Equals(content TemplateIdContent) bool {
	return reflect.DeepEqual(c, content)
}

func (c TemplateIdContent) IsZero() bool {
	return c.Equals(emptyTemplateIdContent)
}

func (c TemplateIdContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TemplateId int                    `json:"templateId"`
		Params     map[string]interface{} `json:"params,omitempty"`
	}{
		TemplateId: c.TemplateId(),
		Params:     c.Params(),
	})
}

type TemplateIdContentOption func(c *TemplateIdContent)

func WithParams(params map[string]interface{}) TemplateIdContentOption {
	return func(c *TemplateIdContent) {
		c.params = params
	}
}
