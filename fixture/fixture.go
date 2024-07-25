// nolint:gomnd // magic numbers
package fixture

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/johnfercher/medium-api/internal/adapters/drivers/rest"
	"github.com/johnfercher/medium-api/internal/core/models"
)

func ModelProduct() *models.Product {
	id, _ := uuid.NewRandom()
	qtd := rand.Intn(10)
	return &models.Product{
		ID:       id.String(),
		Name:     id.String()[:5],
		Type:     id.String()[:3],
		Quantity: qtd,
	}
}

func RestProduct() *rest.Product {
	id, _ := uuid.NewRandom()
	qtd := rand.Intn(10)
	return &rest.Product{
		ID:       id.String(),
		Name:     id.String()[:5],
		Type:     id.String()[:3],
		Quantity: qtd,
	}
}

func RestCreateProduct() *rest.CreateProduct {
	id, _ := uuid.NewRandom()
	qtd := rand.Intn(10)
	return &rest.CreateProduct{
		Name:     id.String()[:5],
		Type:     id.String()[:3],
		Quantity: qtd,
	}
}
