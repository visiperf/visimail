package visimail

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

var (
	ErrEmptyAttachmentFilename = errors.New("attachment filename is empty")
	ErrEmptyAttachmentContent  = errors.New("attachment content is empty")
	ErrInvalidAttachmentURL    = errors.New("attachment url is invalid")
)

type AttachmentType int

const (
	AttachmentTypeContent AttachmentType = iota
	AttachmentTypeURL
)

type Attachment interface {
	Type() AttachmentType
	Validate() error
	IsZero() bool
}

var emptyAttachmentContent = AttachmentContent{}

type AttachmentContent struct {
	filename string
	content  []byte
}

func NewAttachmentContent(filename string, content []byte) AttachmentContent {
	return AttachmentContent{filename, content}
}

func (a AttachmentContent) Filename() string {
	return a.filename
}

func (a AttachmentContent) Content() []byte {
	return a.content
}

func (a AttachmentContent) Base64Content() string {
	return base64.StdEncoding.EncodeToString(a.Content())
}

func (a AttachmentContent) Type() AttachmentType {
	return AttachmentTypeContent
}

func (a AttachmentContent) Equals(attachment AttachmentContent) bool {
	return a.filename == attachment.filename && bytes.Compare(a.content, attachment.content) == 0
}

func (a AttachmentContent) IsZero() bool {
	return a.Equals(emptyAttachmentContent)
}

func (a AttachmentContent) Validate() error {
	if len(a.filename) <= 0 {
		return ErrEmptyAttachmentFilename
	}

	if len(a.content) <= 0 {
		return ErrEmptyAttachmentContent
	}

	return nil
}

func (a AttachmentContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Content string `json:"content"`
		Name    string `json:"name"`
	}{
		Content: a.Base64Content(),
		Name:    a.Filename(),
	})
}

var emptyAttachmentURL = AttachmentURL{}

type AttachmentURL struct {
	url      string
	filename string
}

func NewAttachmentURL(url string, opts ...AttachmentURLOption) AttachmentURL {
	a := AttachmentURL{url: url}
	for _, opt := range opts {
		opt(&a)
	}

	return a
}

func (a AttachmentURL) URL() string {
	return a.url
}

func (a AttachmentURL) Filename() string {
	return a.filename
}

func (a AttachmentURL) Type() AttachmentType {
	return AttachmentTypeURL
}

func (a AttachmentURL) Equals(attachment AttachmentURL) bool {
	return a == attachment
}

func (a AttachmentURL) IsZero() bool {
	return a.Equals(emptyAttachmentURL)
}

func (a AttachmentURL) Validate() error {
	if _, err := url.Parse(a.url); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidAttachmentURL, err)
	}

	return nil
}

func (a AttachmentURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		URL  string `json:"url"`
		Name string `json:"name,omitempty"`
	}{
		URL:  a.URL(),
		Name: a.Filename(),
	})
}

type AttachmentURLOption func(*AttachmentURL)

func WithFilename(name string) AttachmentURLOption {
	return func(a *AttachmentURL) {
		a.filename = name
	}
}

func chunkAttachments(attachments []Attachment, chunkSize int) [][]Attachment {
	var chunks [][]Attachment
	for i := 0; i < len(attachments); i += chunkSize {
		end := i + chunkSize

		if end > len(attachments) {
			end = len(attachments)
		}

		chunks = append(chunks, attachments[i:end])
	}

	return chunks
}
