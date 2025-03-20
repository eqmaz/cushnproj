// Example integration test file

package integration

import (
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test fetching ISA funds
func TestGetIsaFunds(t *testing.T) {
	req, _ := http.NewRequest("GET", BaseURL+"/isa-funds", nil)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err, "HTTP request failed")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")

	log.Println("ISA Funds Response:", string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestGetUserProductAvailable verifies that available products are fetched correctly.
func TestGetUserProductAvailable(t *testing.T) {
	req, _ := http.NewRequest("GET", BaseURL+"/user/product/available", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("userId", "1")

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err, "[TestGetUserProductAvailable] HTTP request failed")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	// Print response
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(body))

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// And so on... with positive and negative test cases
