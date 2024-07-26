// nolint:dupl // separated handler
package grpc

import (
	"context"

	"github.com/johnfercher/medium-api/internal/core/ports"
)

type updateProduct struct {
	UnimplementedUpdateProductHandlerServer
	service ports.ProductService
}

func NewUpdateProduct(service ports.ProductService) *updateProduct {
	return &updateProduct{
		service: service,
	}
}

func (c *updateProduct) Update(ctx context.Context, contract *Product) (*ProductResponse, error) {
	product := MapProductGrpcToModel(contract)

	product, err := c.service.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return MapProductModelToGrpcResponse(product), nil
}
