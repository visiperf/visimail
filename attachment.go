package visimail

type attachmentType int

const (
	attachmentTypeContent attachmentType = iota
	attachmentTypeURL
)

// Attachment is interface implemented by types representing an attachment
type Attachment interface {
	Filename() string
	Type() attachmentType
}

type attachmentContent struct {
	filename string
	content  string
}

// NewAttachmentContent is factory to create a new attachment containing content encoded to base64
func NewAttachmentContent(filename, content string) Attachment {
	return &attachmentContent{filename, content}
}

func (a attachmentContent) Filename() string {
	return a.filename
}

func (a attachmentContent) Content() string {
	return a.content
}

func (a attachmentContent) Type() attachmentType {
	return attachmentTypeContent
}

type attachmentURL struct {
	filename string
	url      string
}

// NewAttachmentURL is factory to create a new attachment containing an external url
func NewAttachmentURL(filename, url string) Attachment {
	return &attachmentURL{filename, url}
}

func (a attachmentURL) Filename() string {
	return a.filename
}

func (a attachmentURL) URL() string {
	return a.url
}

func (a attachmentURL) Type() attachmentType {
	return attachmentTypeURL
}
