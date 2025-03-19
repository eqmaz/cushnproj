package converter

import (
	"database/sql"
	"testing"

	"cushon/internal/models"
	"cushon/internal/router/dto"

	"github.com/stretchr/testify/assert"
)

func TestMapFundsToResponse(t *testing.T) {
	funds := []models.Fund{
		{
			Uuid:        "fund-1",
			Title:       "Fund One",
			Description: sql.NullString{String: "First Fund", Valid: true},
		},
		{
			Uuid:        "fund-2",
			Title:       "Fund Two",
			Description: sql.NullString{String: "", Valid: false}, // Should return empty string
		},
	}

	expected := []dto.FundsResponse{
		{
			Uuid:        "fund-1",
			Title:       "Fund One",
			Description: "First Fund",
		},
		{
			Uuid:        "fund-2",
			Title:       "Fund Two",
			Description: "", // Should be an empty string
		},
	}

	result := MapFundsToResponse(funds)
	assert.Equal(t, expected, result, "Expected mapped funds response to match")
}
