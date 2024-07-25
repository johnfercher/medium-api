package mysql

import (
	"context"
	"strings"

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

func (p *repository) GetByID(_ context.Context, id string) (*models.Product, error) {
	product := &models.Product{}

	tx := p.db.Where("id = ?", id).First(product)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return product, nil
}

func (p *repository) Search(_ context.Context, productType string) ([]*models.Product, error) {
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
		return nil, tx.Error
	}

	return products, nil
}

func (p *repository) Create(_ context.Context, product *models.Product) error {
	tx := p.db.Create(product)
	return tx.Error
}

func (p *repository) Update(_ context.Context, productToUpdate *models.Product) error {
	tx := p.db.Model(&models.Product{}).Where("id = ?", productToUpdate.ID).Updates(map[string]interface{}{
		"name":     productToUpdate.Name,
		"type":     productToUpdate.Type,
		"quantity": productToUpdate.Quantity,
	})

	return tx.Error
}

func (p *repository) Delete(_ context.Context, id string) error {
	tx := p.db.Where("id = ?", id).Delete(&models.Product{})
	return tx.Error
}
