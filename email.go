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

// Body is method to set content of email using body builder
func (b *EmailBuilder) Body() *EmailBodyBuilder {
	return &EmailBodyBuilder{*b}
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
