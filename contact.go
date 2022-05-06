package visimail

import (
	"encoding/json"
	"errors"
	"net/mail"
)

var (
	ErrEmptyContact = errors.New("contact is empty")
)

type Recipient interface {
	json.Marshaler
	Name() string
	Email() string
	Validate() error
}

type Sender interface {
	json.Marshaler
	Name() string
	Email() string
	Validate() error
}

var emptyContact = Contact{}

type Contact struct {
	email string
	name  string
}

func NewContact(email string, opts ...ContactOption) Contact {
	c := Contact{email: email}
	for _, opt := range opts {
		opt(&c)
	}

	return c
}

func (c Contact) Name() string {
	return c.name
}

func (c Contact) Email() string {
	return c.email
}

func (c Contact) Equals(contact Contact) bool {
	return c == contact
}

func (c Contact) IsZero() bool {
	return c.Equals(emptyContact)
}

func (c Contact) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Email string `json:"email"`
		Name  string `json:"name,omitempty"`
	}{
		Email: c.Email(),
		Name:  c.Name(),
	})
}

func (c Contact) Validate() error {
	if c.IsZero() {
		return ErrEmptyContact
	}

	_, err := mail.ParseAddress(c.email)
	return err
}

type ContactOption func(c *Contact)

func WithName(name string) ContactOption {
	return func(c *Contact) {
		c.name = name
	}
}
