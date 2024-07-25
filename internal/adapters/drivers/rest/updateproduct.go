package rest

import (
	"net/http"

	"github.com/johnfercher/medium-api/internal/core/ports"
	"github.com/johnfercher/medium-api/pkg/api/apierror"
	"github.com/johnfercher/medium-api/pkg/api/apiresponse"
)

type updateProduct struct {
	name    string
	verb    string
	pattern string
	service ports.ProductService
}

func NewUpdateProduct(service ports.ProductService) *updateProduct {
	return &updateProduct{
		name:    "update_product",
		pattern: "/products/{id}",
		verb:    "PUT",
		service: service,
	}
}

func (p *updateProduct) Handle(r *http.Request) (apiresponse.APIResponse, apierror.APIError) {
	ctx := r.Context()

	updateProduct, err := DecodeProductFromBodyAndURI(r)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	product := updateProduct.ToProduct()

	updatedProduct, err := p.service.Update(ctx, product)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	response := ProductToResponse(updatedProduct)

	return apiresponse.New(response, http.StatusOK), nil
}

func (p *updateProduct) Name() string {
	return p.name
}

func (p *updateProduct) Pattern() string {
	return p.pattern
}

func (p *updateProduct) Verb() string {
	return p.verb
}
