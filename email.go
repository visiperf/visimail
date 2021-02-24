package visimail

// Email is a struct representing an email
type Email struct {
	from        Sender
	to          []Recipient
	cc          []Recipient
	bcc         []Recipient
	body        string
	bodyType    bodyType
	subject     string
	replyTo     Recipient
	attachments []Attachment
	templateID  int
	params      map[string]interface{}
}

// EmailBuilder should be used to create a new email
type EmailBuilder struct {
	email *Email
}

// NewEmailBuilder is factory to create a new email builder
func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{&Email{}}
}

// From is method to set email sender
func (b *EmailBuilder) From(sender Sender) *EmailBuilder {
	b.email.from = sender

	return b
}

// AppendTo is method to append a recipient in email to list
func (b *EmailBuilder) AppendTo(recipient Recipient) *EmailBuilder {
	b.email.to = append(b.email.to, recipient)

	return b
}

// AppendCC is method to append a recipient in email cc list
func (b *EmailBuilder) AppendCC(recipient Recipient) *EmailBuilder {
	b.email.cc = append(b.email.cc, recipient)

	return b
}

// AppendBCC is method to append a recipient in email bcc list
func (b *EmailBuilder) AppendBCC(recipient Recipient) *EmailBuilder {
	b.email.bcc = append(b.email.bcc, recipient)

	return b
}

// Body is method to set content of email using body builder
func (b *EmailBuilder) Body() *EmailBodyBuilder {
	return &EmailBodyBuilder{*b}
}

// Subject is method to set email subject
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject

	return b
}

// ReplyTo is method to set reply to of email
func (b *EmailBuilder) ReplyTo(recipient Recipient) *EmailBuilder {
	b.email.replyTo = recipient

	return b
}

// EmailBodyBuilder is struct used to set content of email
type EmailBodyBuilder struct {
	EmailBuilder
}

// Text is method to set plain text as email content
func (b *EmailBodyBuilder) Text(text string) *EmailBodyBuilder {
	b.email.body = text
	b.email.bodyType = bodyTypeText

	return b
}

// HTML is method to set html as email content
func (b *EmailBodyBuilder) HTML(html string) *EmailBodyBuilder {
	b.email.body = html
	b.email.bodyType = bodyTypeHTML

	return b
}
