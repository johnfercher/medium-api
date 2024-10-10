package services

import (
	"context"

	"github.com/johnfercher/medium-api/pkg/observability/log"
	"github.com/johnfercher/medium-api/pkg/observability/log/field"

	"github.com/google/uuid"
	"github.com/johnfercher/medium-api/internal/core/models"
	"github.com/johnfercher/medium-api/internal/core/ports"
)

type ProductService struct {
	productRepository ports.ProductRepository
}

func New(productRepository ports.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (p *ProductService) GetByID(ctx context.Context, id string) (*models.Product, error) {
	return p.productRepository.GetByID(ctx, id)
}

func (p *ProductService) Search(ctx context.Context, productType string) ([]*models.Product, error) {
	return p.productRepository.Search(ctx, productType)
}

func (p *ProductService) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Error(ctx, "could not create product id", field.Error(err))
		return nil, err
	}

	idString := id.String()
	product.ID = idString

	err = p.productRepository.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) Update(ctx context.Context, product *models.Product) (*models.Product, error) {
	err := p.productRepository.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) Delete(ctx context.Context, id string) error {
	return p.productRepository.Delete(ctx, id)
}
