package visimail

import "fmt"

type HttpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (he HttpError) Error() string {
	return fmt.Sprintf("%s: %s", he.Code, he.Message)
}

func isHttpError(statusCode int) bool {
	return statusCode < 200 || statusCode >= 300
}
