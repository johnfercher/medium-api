package ports

import (
	"context"

	"github.com/johnfercher/medium-api/internal/core/models"
)

type ProductService interface {
	GetByID(ctx context.Context, id string) (*models.Product, error)
	Search(ctx context.Context, productType string) ([]*models.Product, error)
	Create(ctx context.Context, product *models.Product) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) (*models.Product, error)
	Delete(ctx context.Context, id string) error
}
