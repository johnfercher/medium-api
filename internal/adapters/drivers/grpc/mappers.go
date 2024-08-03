package grpc

import "github.com/johnfercher/medium-api/internal/core/models"

func MapCreateProductToModel(createProduct *CreateProduct) *models.Product {
	return &models.Product{
		Name:     createProduct.GetName(),
		Type:     createProduct.GetType(),
		Quantity: int(createProduct.GetQuantity()),
	}
}

func MapProductModelToGrpc(product *models.Product) *Product {
	return &Product{
		Id:       product.ID,
		Name:     product.Name,
		Type:     product.Type,
		Quantity: int64(product.Quantity),
	}
}

func MapProductGrpcToModel(product *Product) *models.Product {
	return &models.Product{
		ID:       product.GetId(),
		Name:     product.GetName(),
		Type:     product.GetType(),
		Quantity: int(product.GetQuantity()),
	}
}

func MapProductModelToGrpcResponse(product *models.Product) *ProductResponse {
	return &ProductResponse{
		Product: MapProductModelToGrpc(product),
	}
}

func MapProductsModelToGrpcResponse(models []*models.Product) *ProductsResponse {
	products := &ProductsResponse{}

	for _, model := range models {
		products.Products = append(products.Products, MapProductModelToGrpc(model))
	}

	return products
}
