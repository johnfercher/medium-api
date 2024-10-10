package apierror

import "fmt"

type APIError interface {
	Name() string
	Code() int
	Error() string
}

type apiError struct {
	name string
	code int
}

func New(name string, code int) APIError {
	return &apiError{name, code}
}

func (a *apiError) Name() string {
	return a.name
}

func (a *apiError) Code() int {
	return a.code
}

func (a *apiError) Error() string {
	return fmt.Sprintf("%s (%d)", a.name, a.code)
}
