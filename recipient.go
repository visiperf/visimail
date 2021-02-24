package visimail

// Recipient is interface implemented by types that represents a person receiving an email
type Recipient interface {
	Name() string
	Email() string
}
