// nolint:dupl // separated handler
package grpc

import (
	"context"

	"github.com/johnfercher/medium-api/internal/core/ports"
)

type getProductByID struct {
	UnimplementedGetProductHandlerServer
	service ports.ProductService
}

func NewGetProductByID(service ports.ProductService) *getProductByID {
	return &getProductByID{
		service: service,
	}
}

func (c *getProductByID) Get(ctx context.Context, id *ID) (*ProductResponse, error) {
	product, err := c.service.GetByID(ctx, id.GetId())
	if err != nil {
		return nil, err
	}

	return MapProductModelToGrpcResponse(product), nil
}
