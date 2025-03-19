package converter

import (
	"database/sql"
	"testing"

	"cushon/internal/models"
	"cushon/internal/router/dto"

	"github.com/stretchr/testify/assert"
)

func TestMapProductsToResponse(t *testing.T) {
	products := []models.Product{
		{
			Uuid:        "product-1",
			Title:       "Product One",
			Type:        models.ProductTypeDirect,
			Description: sql.NullString{String: "First Product", Valid: true},
		},
		{
			Uuid:        "product-2",
			Title:       "Product Two",
			Type:        models.ProductTypeEmployer,
			Description: sql.NullString{String: "", Valid: false}, // Should return empty string
		},
	}

	expected := []dto.ProductResponse{
		{
			Uuid:        "product-1",
			Title:       "Product One",
			Type:        models.ProductTypeDirect,
			Description: "First Product",
		},
		{
			Uuid:        "product-2",
			Title:       "Product Two",
			Type:        models.ProductTypeEmployer,
			Description: "", // Should be an empty string
		},
	}

	result := MapProductsToResponse(products)
	assert.Equal(t, expected, result, "Expected mapped products response to match")
}
