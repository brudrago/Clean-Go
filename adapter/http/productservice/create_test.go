package productService

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brudrago/clean-go/core/domain"
	"github.com/brudrago/clean-go/core/domain/mocks"
	"github.com/brudrago/clean-go/core/dto"
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupCreate(t *testing.T) (dto.CreateProductRequest, domain.Product, *gomock.Controller) {
	fakeProductRequest := dto.CreateProductRequest{}
	fakeProduct := domain.Product{}

	_ = faker.FakeData(&fakeProductRequest)
	_ = faker.FakeData(&fakeProduct)

	if fakeProduct.ID == 0 {
		fakeProduct.ID = 1
	}

	mockCtrl := gomock.NewController(t)

	return fakeProductRequest, fakeProduct, mockCtrl
}

func TestCreate(t *testing.T) {
	fakeProductRequest, fakeProduct, mock := setupCreate(t)
	defer mock.Finish()

	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().
		Create(&fakeProductRequest).
		Return(&fakeProduct, nil)

	sut := New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")

	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCreate_JsonErrorFormater(t *testing.T) {
	_, _, mock := setupCreate(t)
	defer mock.Finish()

	mockProductUseCase := mocks.NewMockProductUseCase(mock)

	sut := New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader("{")) // JSON inválido
	r.Header.Set("Content-Type", "application/json")

	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	// Espera qualquer coisa diferente de 200
	require.NotEqual(t, http.StatusOK, res.StatusCode)
}

func TestCreate_ProductError(t *testing.T) {
	fakeProductRequest, _, mock := setupCreate(t)
	defer mock.Finish()

	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().
		Create(&fakeProductRequest).
		Return(nil, fmt.Errorf("ANY ERROR"))

	sut := New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")

	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.NotEqual(t, http.StatusOK, res.StatusCode)
}
