package dto

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromJSONCreateProductRequest(t *testing.T) {
	jsonBody := `{
		"name": "Product Test",
		"price": 99.99,
		"description": "This is a test product"
	}`

	createProductRequest, err := FromJSONCreateProductRequest(strings.NewReader(jsonBody))

	require.NoError(t, err)
	require.Equal(t, "Product Test", createProductRequest.Name)
	require.Equal(t, 99.99, createProductRequest.Price)
	require.Equal(t, "This is a test product", createProductRequest.Description)
}

func TestFromJSONCreateProductRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := FromJSONCreateProductRequest(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}
