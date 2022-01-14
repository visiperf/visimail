package visimail

type AttachmentType int

const (
	AttachmentTypeContent AttachmentType = iota
	AttachmentTypeURL
)

// Attachment is interface implemented by types representing an attachment
type Attachment interface {
	Filename() string
	Type() AttachmentType
}

type AttachmentContent struct {
	filename string
	content  string
}

// NewAttachmentContent is factory to create a new attachment containing content encoded to base64
func NewAttachmentContent(filename, content string) Attachment {
	return &AttachmentContent{filename, content}
}

func (a AttachmentContent) Filename() string {
	return a.filename
}

func (a AttachmentContent) Content() string {
	return a.content
}

func (a AttachmentContent) Type() AttachmentType {
	return AttachmentTypeContent
}

type AttachmentURL struct {
	filename string
	url      string
}

// NewAttachmentURL is factory to create a new attachment containing an external url
func NewAttachmentURL(filename, url string) Attachment {
	return &AttachmentURL{filename, url}
}

func (a AttachmentURL) Filename() string {
	return a.filename
}

func (a AttachmentURL) URL() string {
	return a.url
}

func (a AttachmentURL) Type() AttachmentType {
	return AttachmentTypeURL
}
