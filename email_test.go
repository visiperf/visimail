package visimail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmaiValidate(t *testing.T) {
	email := Email{
		from: NewContact("support@visiperf.io", WithName("Visiperf")),
		to: []Recipient{
			NewContact("john.doe@visiperf.com", WithName("John Doe")),
		},
		cc: []Recipient{
			NewContact("elaine.eastwood@visiperf.com", WithName("Elaine Eastwood")),
		},
		bcc: []Recipient{
			NewContact("simon.york@visiperf.com", WithName("Simon York")),
		},
		body: NewTemplateContent(1, WithParams(map[string]interface{}{
			"name": "John Doe",
		})),
		subject: "Hello",
		replyTo: NewContact("support@visiperf.io", WithName("Visiperf")),
		attachments: []Attachment{
			NewAttachmentURL("https://visiretail-dev.com/static/media/visiretail-logo.9ba46052.svg", WithFilename("visiretail-logo.svg")),
		},
	}

	tests := []struct {
		name  string
		email func() *Email
		err   error
	}{{
		name: "text content without from sender",
		email: func() *Email {
			e := email.copy()

			e.body = NewPlainTextContent("hello")
			e.from = nil

			return e
		},
		err: ErrFromRequired,
	}, {
		name: "no recipients in to",
		email: func() *Email {
			e := email.copy()

			e.to = nil

			return e
		},
		err: ErrToRequired,
	}, {
		name: "no body",
		email: func() *Email {
			e := email.copy()

			e.body = nil

			return e
		},
		err: ErrBodyRequired,
	}, {
		name: "no reply to",
		email: func() *Email {
			e := email.copy()

			e.replyTo = nil

			return e
		},
		err: ErrReplyToRequired,
	}, {
		name: "text content without subject",
		email: func() *Email {
			e := email.copy()

			e.body = NewPlainTextContent("hello")
			e.subject = ""

			return e
		},
		err: ErrSubjectRequired,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.err, test.email().Validate())
		})
	}
}
