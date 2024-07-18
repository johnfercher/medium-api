package services_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/johnfercher/medium-api/internal/core/models"
	"github.com/johnfercher/medium-api/internal/services"
	"github.com/johnfercher/medium-api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := services.New(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*services.ProductService", fmt.Sprintf("%T", sut))
}

func TestProductService_Create_WhenCannotCreate_ShouldReturnAnError(t *testing.T) {
	// Arrange
	ctx := context.TODO()
	productToCreate := &models.Product{
		Name:     "name",
		Type:     "type",
		Quantity: 100,
	}

	errToReturn := errors.New("any_error")

	repository := mocks.NewProductRepository(t)
	repository.On("Create", ctx, productToCreate).Return(errToReturn)

	sut := services.New(repository)

	// Act
	product, err := sut.Create(ctx, productToCreate)

	// Assert
	assert.Nil(t, product)
	assert.Equal(t, errToReturn, err)
	repository.AssertNumberOfCalls(t, "Create", 1)
	repository.AssertCalled(t, "Create", ctx, productToCreate)
}

func TestProductService_Create_WhenEverythingWorks_ShouldReturnProduct(t *testing.T) {
	// Arrange
	ctx := context.TODO()
	productToCreate := &models.Product{
		Name:     "name",
		Type:     "type",
		Quantity: 100,
	}

	repository := mocks.NewProductRepository(t)
	repository.On("Create", ctx, productToCreate).Return(nil)

	sut := services.New(repository)

	// Act
	product, err := sut.Create(ctx, productToCreate)

	// Assert
	assert.Equal(t, productToCreate, product)
	assert.NotEmpty(t, product.ID)
	assert.Nil(t, err)
	repository.AssertNumberOfCalls(t, "Create", 1)
	repository.AssertCalled(t, "Create", ctx, productToCreate)
}
