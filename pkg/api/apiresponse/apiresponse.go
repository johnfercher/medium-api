package apiresponse

type APIResponse interface {
	Object() interface{}
	Code() int
}

type apiResponse struct {
	object interface{}
	code   int
}

func New(object interface{}, code int) APIResponse {
	return &apiResponse{object, code}
}

func (a *apiResponse) Object() interface{} {
	return a.object
}

func (a *apiResponse) Code() int {
	return a.code
}
