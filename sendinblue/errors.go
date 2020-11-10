package sendinblue

import "fmt"

// ParamsError is error occured when invalid parameters is specified
type ParamsError struct {
	Err error
}

func (e *ParamsError) Error() string {
	return e.Err.Error()
}

// QueryError is error occured when sendinblue api call fail
type QueryError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
