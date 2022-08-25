package visimail

import "errors"

var (
	ErrEmptyTag = errors.New("tag is empty")
)

type Tag string

func (t Tag) IsEmpty() bool {
	return t == ""
}

func (t Tag) String() string {
	return string(t)
}
