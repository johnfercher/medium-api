// nolint:dupl // separated handler
package rest

import (
	"net/http"

	"github.com/johnfercher/medium-api/internal/core/ports"
	"github.com/johnfercher/medium-api/pkg/api/apierror"
	"github.com/johnfercher/medium-api/pkg/api/apiresponse"
)

type getProductByID struct {
	name    string
	verb    string
	pattern string
	service ports.ProductService
}

func NewGetProductByID(service ports.ProductService) *getProductByID {
	return &getProductByID{
		name:    "get_product_by_id",
		pattern: "/products/{id}",
		verb:    "GET",
		service: service,
	}
}

func (p *getProductByID) Handle(r *http.Request) (apiresponse.APIResponse, apierror.APIError) {
	ctx := r.Context()

	id, err := DecodeStringIDFromURI(r)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	product, err := p.service.GetByID(ctx, id)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	response := ProductToResponse(product)

	return apiresponse.New(response, http.StatusOK), nil
}

func (p *getProductByID) Name() string {
	return p.name
}

func (p *getProductByID) Pattern() string {
	return p.pattern
}

func (p *getProductByID) Verb() string {
	return p.verb
}
