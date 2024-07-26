// nolint:dupl // separated handler
package grpc

import (
	"context"

	"github.com/johnfercher/medium-api/internal/core/ports"
)

type createProduct struct {
	UnimplementedCreateProductHandlerServer
	service ports.ProductService
}

func NewCreateProduct(service ports.ProductService) *createProduct {
	return &createProduct{
		service: service,
	}
}

func (c *createProduct) Create(ctx context.Context, createProduct *CreateProduct) (*ProductResponse, error) {
	product := MapCreateProductToModel(createProduct)

	product, err := c.service.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return MapProductModelToGrpcResponse(product), nil
}
