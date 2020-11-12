package visimail

// EmailService represent available methods of service used to send emails
type EmailService interface {
	SendEmailFromTemplate(templateID string, subject *string, to []string, cc []string, bcc []string, params map[string]interface{}, attachments []*Attachment, replyTo string) (string, error)
}
