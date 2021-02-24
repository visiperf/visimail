package visimail

// EmailSender is interface implemented by types that represents a service providing email sending
type EmailSender interface {
	SendEmail(email *Email, opts ...FuncOption) ([]string, error)
}

type options struct {
	nbAttachmentsPerChunk int
}

// FuncOption is type used to apply some values to options
type FuncOption func(*options)

// WithChunkedAttachments is a func option used to chunk attachments in multiple parts
func WithChunkedAttachments(nbAttachmentsPerChunk int) FuncOption {
	return func(o *options) {
		o.nbAttachmentsPerChunk = nbAttachmentsPerChunk
	}
}

func buildOptions(opts ...FuncOption) options {
	var os options
	for _, opt := range opts {
		opt(&os)
	}

	return os
}
