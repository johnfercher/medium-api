// nolint:dupl // separated handler
package grpc

import (
	"context"

	"github.com/johnfercher/medium-api/internal/core/ports"
)

type deleteProduct struct {
	UnimplementedDeleteProductHandlerServer
	service ports.ProductService
}

func NewDeleteProduct(service ports.ProductService) *deleteProduct {
	return &deleteProduct{
		service: service,
	}
}

func (c *deleteProduct) Delete(ctx context.Context, id *ID) (*Empty, error) {
	err := c.service.Delete(ctx, id.GetId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}
