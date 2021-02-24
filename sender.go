package visimail

// Sender is interface implemented by types that represents a person sending a email
type Sender interface {
	Name() string
	Email() string
}
