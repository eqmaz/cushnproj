package converter

import (
	"cushon/internal/models"
	"cushon/internal/router/dto"
)

// MapFundsToResponse - ensure we only return the necessary fields
func MapFundsToResponse(products []models.Fund) []dto.FundsResponse {
	result := make([]dto.FundsResponse, len(products))
	for i, p := range products {
		result[i] = dto.FundsResponse{
			Uuid:        p.Uuid,
			Title:       p.Title,
			Description: p.Description.String, // Ensure it's always a string
		}
	}
	return result
}
