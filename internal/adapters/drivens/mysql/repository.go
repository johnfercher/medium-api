package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/johnfercher/medium-api/pkg/observability/log"
	"github.com/johnfercher/medium-api/pkg/observability/log/field"

	"github.com/johnfercher/medium-api/internal/core/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (p *repository) GetByID(ctx context.Context, id string) (*models.Product, error) {
	product := &models.Product{}

	tx := p.db.Where("id = ?", id).First(product)

	if tx.Error != nil {
		log.Error(ctx, fmt.Sprintf("could not found product %s", id), field.Error(tx.Error))
		return nil, tx.Error
	}

	return product, nil
}

func (p *repository) Search(ctx context.Context, productType string) ([]*models.Product, error) {
	limit := 100
	products := []*models.Product{}

	query := []string{}
	args := []interface{}{}

	query = append(query, "products.type = ?")
	args = append(args, productType)

	tx := p.db.Table("products").
		Select("products.id, products.name, products.type, products.quantity")

	tx = tx.Where(strings.Join(query, " AND "), args...)

	tx = tx.Limit(limit)

	tx = tx.Scan(&products)

	if tx.Error != nil {
		log.Error(ctx, fmt.Sprintf("could not search %s", productType), field.Error(tx.Error))
		return nil, tx.Error
	}

	return products, nil
}

func (p *repository) Create(ctx context.Context, product *models.Product) error {
	tx := p.db.Create(product)
	if tx.Error != nil {
		log.Error(ctx, fmt.Sprintf("could not create product %v", product), field.Error(tx.Error))
		return tx.Error
	}

	return nil
}

func (p *repository) Update(ctx context.Context, productToUpdate *models.Product) error {
	tx := p.db.Model(&models.Product{}).Where("id = ?", productToUpdate.ID).Updates(map[string]interface{}{
		"name":     productToUpdate.Name,
		"type":     productToUpdate.Type,
		"quantity": productToUpdate.Quantity,
	})
	if tx.Error != nil {
		log.Error(ctx, fmt.Sprintf("could not update product %v", productToUpdate), field.Error(tx.Error))
		return tx.Error
	}

	return nil
}

func (p *repository) Delete(ctx context.Context, id string) error {
	tx := p.db.Where("id = ?", id).Delete(&models.Product{})
	if tx.Error != nil {
		log.Error(ctx, fmt.Sprintf("could not update product %s", id), field.Error(tx.Error))
		return tx.Error
	}

	return nil
}
