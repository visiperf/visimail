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
