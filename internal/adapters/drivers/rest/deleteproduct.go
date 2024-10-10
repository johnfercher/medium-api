package rest

import (
	"net/http"

	"github.com/johnfercher/medium-api/pkg/observability/log"
	"github.com/johnfercher/medium-api/pkg/observability/log/field"

	"github.com/johnfercher/medium-api/internal/core/ports"
	"github.com/johnfercher/medium-api/pkg/api/apierror"
	"github.com/johnfercher/medium-api/pkg/api/apiresponse"
)

type deleteProduct struct {
	name    string
	verb    string
	pattern string
	service ports.ProductService
}

func NewDeleteProduct(service ports.ProductService) *deleteProduct {
	return &deleteProduct{
		name:    "delete_product",
		pattern: "/products/{id}",
		verb:    "DELETE",
		service: service,
	}
}

func (p *deleteProduct) Handle(r *http.Request) (apiresponse.APIResponse, apierror.APIError) {
	ctx := r.Context()

	id, err := DecodeStringIDFromURI(r)
	if err != nil {
		log.Error(ctx, "could not get id from uri", field.Error(err),
			field.StatusCode(http.StatusBadRequest))
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	err = p.service.Delete(ctx, id)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	return apiresponse.New(nil, http.StatusNoContent), nil
}

func (p *deleteProduct) Name() string {
	return p.name
}

func (p *deleteProduct) Pattern() string {
	return p.pattern
}

func (p *deleteProduct) Verb() string {
	return p.verb
}
