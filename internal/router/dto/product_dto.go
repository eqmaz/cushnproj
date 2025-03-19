package dto

import "cushon/internal/models"

type ProductResponse struct {
	Uuid        string             `json:"uuid"`
	Title       string             `json:"title"`
	Type        models.ProductType `json:"type"`
	Description string             `json:"description,omitempty"`
}
