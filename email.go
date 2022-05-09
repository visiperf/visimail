package visimail

import (
	"encoding/json"
	"errors"
)

var (
	ErrFromRequired    = errors.New("from is required")
	ErrToRequired      = errors.New("to is required")
	ErrBodyRequired    = errors.New("body is required")
	ErrReplyToRequired = errors.New("reply to is required")
	ErrSubjectRequired = errors.New("subject is required")
)

type Email struct {
	from        Sender
	to          []Recipient
	cc          []Recipient
	bcc         []Recipient
	body        Content
	subject     string
	replyTo     Recipient
	attachments []Attachment
}

func (e *Email) Validate() error {
	if _, ok := e.body.(TemplateContent); !ok {
		if e.from == nil {
			return ErrFromRequired
		}
	}

	if e.from != nil {
		if err := e.from.Validate(); err != nil {
			return err
		}
	}

	if len(e.to) <= 0 {
		return ErrToRequired
	}

	for _, r := range e.to {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	for _, r := range e.cc {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	for _, r := range e.bcc {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	if e.body == nil {
		return ErrBodyRequired
	}

	if err := e.body.Validate(); err != nil {
		return err
	}

	if e.replyTo == nil {
		return ErrReplyToRequired
	}

	if err := e.replyTo.Validate(); err != nil {
		return err
	}

	for _, a := range e.attachments {
		if err := a.Validate(); err != nil {
			return err
		}
	}

	if len(e.subject) <= 0 {
		return ErrSubjectRequired
	}

	return nil
}

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Content
		Sender      Sender       `json:"sender,omitempty"`
		To          []Recipient  `json:"to"`
		CC          []Recipient  `json:"cc"`
		BCC         []Recipient  `json:"bcc"`
		Subject     string       `json:"subject"`
		ReplyTo     Sender       `json:"replyTo"`
		Attachments []Attachment `json:"attachments"`
	}{
		Content:     e.body,
		Sender:      e.from,
		To:          e.to,
		CC:          e.cc,
		BCC:         e.bcc,
		Subject:     e.subject,
		ReplyTo:     e.replyTo,
		Attachments: e.attachments,
	})
}

type EmailBuilder struct {
	email *Email
}

func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{&Email{}}
}

func (b *EmailBuilder) From(sender Sender) *EmailBuilder {
	b.email.from = sender

	return b
}

func (b *EmailBuilder) AppendTo(recipient Recipient) *EmailBuilder {
	b.email.to = append(b.email.to, recipient)

	return b
}

func (b *EmailBuilder) AppendCC(recipient Recipient) *EmailBuilder {
	b.email.cc = append(b.email.cc, recipient)

	return b
}

func (b *EmailBuilder) AppendBCC(recipient Recipient) *EmailBuilder {
	b.email.bcc = append(b.email.bcc, recipient)

	return b
}

func (b *EmailBuilder) Body() *EmailBodyBuilder {
	return &EmailBodyBuilder{*b}
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject

	return b
}

func (b *EmailBuilder) ReplyTo(recipient Recipient) *EmailBuilder {
	b.email.replyTo = recipient

	return b
}

func (b *EmailBuilder) AppendAttachment(attachment Attachment) *EmailBuilder {
	b.email.attachments = append(b.email.attachments, attachment)

	return b
}

type EmailBodyBuilder struct {
	EmailBuilder
}

func (b *EmailBodyBuilder) HTML(html string) *EmailBodyBuilder {
	b.email.body = NewHTMLContent(html)

	return b
}

func (b *EmailBodyBuilder) PlainText(text string) *EmailBodyBuilder {
	b.email.body = NewPlainTextContent(text)

	return b
}

func (b *EmailBodyBuilder) Template(id int, opts ...TemplateContentOption) *EmailBodyBuilder {
	b.email.body = NewTemplateContent(id, opts...)

	return b
}
