package dto

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromValuePaginationRequestParams(t *testing.T) {
	fakeRequest := httptest.NewRequest(http.MethodGet, "/product", nil)
	queryStringParams := fakeRequest.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")
	fakeRequest.URL.RawQuery = queryStringParams.Encode()

	// aqui chama direto, sem dto.
	paginationRequest, err := FromValuePaginationRequestParams(fakeRequest)

	require.NoError(t, err)
	require.Equal(t, 1, paginationRequest.Page)
	require.Equal(t, 10, paginationRequest.ItemsPerPage)
	require.Equal(t, []string{""}, paginationRequest.Sort)
	require.Equal(t, []string{""}, paginationRequest.Descending)
	require.Equal(t, "", paginationRequest.Search)
}
