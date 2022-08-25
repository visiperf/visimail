package visimail

type Tag string

func (t Tag) IsEmpty() bool {
	return t == ""
}

func (t Tag) String() string {
	return string(t)
}
