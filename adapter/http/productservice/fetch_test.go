package productService

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brudrago/clean-go/core/domain"
	"github.com/brudrago/clean-go/core/domain/mocks"
	"github.com/brudrago/clean-go/core/dto"
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupFetch(t *testing.T) (dto.PaginationRequestParams, domain.Product, *gomock.Controller) {
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         []string{""},
		Descending:   []string{""},
		Search:       "",
	}

	fakeProduct := domain.Product{}
	_ = faker.FakeData(&fakeProduct)
	if fakeProduct.ID == 0 {
		fakeProduct.ID = 1
	}

	mockCtrl := gomock.NewController(t)

	return fakePaginationRequestParams, fakeProduct, mockCtrl
}

func TestFetch(t *testing.T) {
	fakePaginationRequestParams, fakeProduct, mock := setupFetch(t)
	defer mock.Finish()

	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.
		EXPECT().
		Fetch(&fakePaginationRequestParams).
		Return(&domain.Pagination[[]domain.Product]{
			Items: []domain.Product{fakeProduct},
			Total: 1,
		}, nil)

	// mesmo pacote, usa New direto
	sut := New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/product", nil)
	r.Header.Set("Content-Type", "application/json")

	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")
	r.URL.RawQuery = queryStringParams.Encode()

	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestFetch_ProductError(t *testing.T) {
	fakePaginationRequestParams, _, mock := setupFetch(t)
	defer mock.Finish()

	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.
		EXPECT().
		Fetch(&fakePaginationRequestParams).
		Return(nil, fmt.Errorf("ANY ERROR"))

	sut := New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/product", nil)
	r.Header.Set("Content-Type", "application/json")

	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")
	r.URL.RawQuery = queryStringParams.Encode()

	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	// qualquer coisa diferente de 200 já indica erro tratado
	require.NotEqual(t, http.StatusOK, res.StatusCode)
}
