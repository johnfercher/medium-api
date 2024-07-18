package api

import (
	"net/http"

	"github.com/johnfercher/medium-api/pkg/api/apierror"
	"github.com/johnfercher/medium-api/pkg/api/apiresponse"
)

type HTTPHandler interface {
	Name() string
	Pattern() string
	Verb() string
	Handle(r *http.Request) (response apiresponse.APIResponse, err apierror.APIError)
}
