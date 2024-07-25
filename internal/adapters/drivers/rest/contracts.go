package rest

import "github.com/johnfercher/medium-api/internal/core/models"

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

func (c *Product) ToProduct() *models.Product {
	return &models.Product{
		ID:       c.ID,
		Name:     c.Name,
		Type:     c.Type,
		Quantity: c.Quantity,
	}
}

func ProductsToResponse(models []*models.Product) []*Product {
	var response []*Product
	for _, model := range models {
		response = append(response, ProductToResponse(model))
	}
	return response
}

func ProductToResponse(model *models.Product) *Product {
	return &Product{
		ID:       model.ID,
		Name:     model.Name,
		Type:     model.Type,
		Quantity: model.Quantity,
	}
}

type CreateProduct struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

func (c *CreateProduct) ToProduct() *models.Product {
	return &models.Product{
		Name:     c.Name,
		Type:     c.Type,
		Quantity: c.Quantity,
	}
}
