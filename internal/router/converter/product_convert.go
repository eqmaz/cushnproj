package converter

import (
	"cushon/internal/models"
	"cushon/internal/router/dto"
)

// MapProductsToResponse - ensure we only return the necessary fields
// We don't expose database primary key IDs to the outside world
// Doing this explicitly ensures we don't accidentally leak sensitive data
func MapProductsToResponse(products []models.Product) []dto.ProductResponse {
	responseProducts := make([]dto.ProductResponse, len(products))
	for i, p := range products {
		responseProducts[i] = dto.ProductResponse{
			Uuid:        p.Uuid,
			Title:       p.Title,
			Type:        p.Type,
			Description: p.Description.String, // Ensure it's always a string
		}
	}
	return responseProducts
}
