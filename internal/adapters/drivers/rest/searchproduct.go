package rest

import (
	"net/http"

	"github.com/johnfercher/medium-api/internal/core/ports"
	"github.com/johnfercher/medium-api/pkg/api/apierror"
	"github.com/johnfercher/medium-api/pkg/api/apiresponse"
)

type searchProducts struct {
	name    string
	verb    string
	pattern string
	service ports.ProductService
}

func NewSearchProducts(service ports.ProductService) *searchProducts {
	return &searchProducts{
		name:    "search_product",
		pattern: "/products",
		verb:    "GET",
		service: service,
	}
}

func (p *searchProducts) Handle(r *http.Request) (apiresponse.APIResponse, apierror.APIError) {
	ctx := r.Context()

	productType := DecodeTypeQueryString(r)

	products, err := p.service.Search(ctx, productType)
	if err != nil {
		return nil, apierror.New(err.Error(), http.StatusBadRequest)
	}

	response := ProductsToResponse(products)

	return apiresponse.New(response, http.StatusOK), nil
}

func (p *searchProducts) Name() string {
	return p.name
}

func (p *searchProducts) Pattern() string {
	return p.pattern
}

func (p *searchProducts) Verb() string {
	return p.verb
}
