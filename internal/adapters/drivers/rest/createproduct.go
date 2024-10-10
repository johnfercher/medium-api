// nolint:dupl // separated handler
package rest

import (
	"net/http"

	"github.com/johnfercher/medium-api/pkg/observability/log"
	"github.com/johnfercher/medium-api/pkg/observability/log/field"

	"github.com/johnfercher/medium-api/internal/core/ports"
	"github.com/johnfercher/medium-api/pkg/api/apierror"
	"github.com/johnfercher/medium-api/pkg/api/apiresponse"
)

type createProduct struct {
	name    string
	verb    string
	pattern string
	service ports.ProductService
}

func NewCreateProduct(service ports.ProductService) *createProduct {
	return &createProduct{
		name:    "create_product",
		pattern: "/products",
		verb:    "POST",
		service: service,
	}
}

func (p *createProduct) Handle(r *http.Request) (apiresponse.APIResponse, apierror.APIError) {
	ctx := r.Context()

	createProduct, err := DecodeCreateProductFromBody(r)
	if err != nil {
		log.Error(ctx, "could not decode create product", field.Error(err),
			field.StatusCode(http.StatusBadRequest))
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	product := createProduct.ToProduct()

	createdProduct, err := p.service.Create(ctx, product)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	return apiresponse.New(createdProduct, http.StatusCreated), nil
}

func (p *createProduct) Name() string {
	return p.name
}

func (p *createProduct) Pattern() string {
	return p.pattern
}

func (p *createProduct) Verb() string {
	return p.verb
}
