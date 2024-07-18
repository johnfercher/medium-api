package rest_test

import (
	"testing"

	"github.com/johnfercher/medium-api/pkg/fixture"
	"github.com/stretchr/testify/assert"
)

func TestProduct_ToProduct(t *testing.T) {
	t.Run("when map to product, then map correctly", func(t *testing.T) {
		// Arrange
		contract := fixture.RestProduct()

		// Act
		model := contract.ToProduct()

		// Assert
		assert.Equal(t, contract.ID, model.ID)
		assert.Equal(t, contract.Name, model.Name)
		assert.Equal(t, contract.Type, model.Type)
		assert.Equal(t, contract.Quantity, model.Quantity)
	})
}
