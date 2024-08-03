// nolint:dupl // separated handler
package grpc

import (
	"context"

	"github.com/johnfercher/medium-api/internal/core/ports"
)

type searchProduct struct {
	UnimplementedSearchProductHandlerServer
	service ports.ProductService
}

func NewSearchProduct(service ports.ProductService) *searchProduct {
	return &searchProduct{
		service: service,
	}
}

func (c *searchProduct) Search(ctx context.Context, t *Type) (*ProductsResponse, error) {
	products, err := c.service.Search(ctx, t.GetType())
	if err != nil {
		return nil, err
	}

	return MapProductsModelToGrpcResponse(products), nil
}
