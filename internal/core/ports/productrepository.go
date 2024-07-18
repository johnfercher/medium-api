package ports

import (
	"context"

	"github.com/johnfercher/medium-api/internal/core/models"
)

type ProductRepository interface {
	GetByID(_ context.Context, id string) (*models.Product, error)
	Search(_ context.Context, productType string) ([]*models.Product, error)
	Create(_ context.Context, product *models.Product) error
	Update(_ context.Context, productToUpdate *models.Product) error
	Delete(_ context.Context, id string) error
}
