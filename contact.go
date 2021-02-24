package visimail

// Recipient is interface implemented by types that represents a person receiving an email
type Recipient interface {
	Name() string
	Email() string
}

// Sender is interface implemented by types that represents a person sending an email
type Sender interface {
	Name() string
	Email() string
}

// Contact is struct representing a person that can be contacted by email
type Contact struct {
	name  string
	email string
}

// NewContact is factory to create a new contact
func NewContact(name, email string) *Contact {
	return &Contact{name, email}
}

// Name return the contact name
func (c Contact) Name() string {
	return c.name
}

// Email return the email address of the contact
func (c Contact) Email() string {
	return c.email
}
