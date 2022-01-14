package visimail

import "encoding/json"

type Recipient interface {
	Name() string
	Email() string
	IsZero() bool
}

type Sender interface {
	Name() string
	Email() string
	IsZero() bool
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

type ContactOption func(c *Contact)

func WithName(name string) ContactOption {
	return func(c *Contact) {
		c.name = name
	}
}
