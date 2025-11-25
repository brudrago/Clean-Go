package productUseCase

import (
	"fmt"
	"testing"

	"github.com/brudrago/clean-go/core/domain"
	"github.com/brudrago/clean-go/core/domain/mocks"
	"github.com/brudrago/clean-go/core/dto"
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}

	fakeDBProduct := domain.Product{}
	_ = faker.FakeData(&fakeDBProduct)

	// garante que o ID não seja zero, pra não quebrar o NotEmpty
	if fakeDBProduct.ID == 0 {
		fakeDBProduct.ID = 1
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.
		EXPECT().
		Fetch(&fakePaginationRequestParams).
		Return(&domain.Pagination[[]domain.Product]{
			Items: []domain.Product{fakeDBProduct},
			Total: 1,
		}, nil)

	sut := New(mockProductRepository)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.NoError(t, err)
	require.NotNil(t, products)
	require.Equal(t, int32(1), products.Total)

	// products.Items já é []domain.Product
	for _, product := range products.Items {
		require.NotEmpty(t, product.ID)
		require.Equal(t, fakeDBProduct.Name, product.Name)
		require.Equal(t, fakeDBProduct.Price, product.Price)
		require.Equal(t, fakeDBProduct.Description, product.Description)
	}
}

func TestFetch_Error(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.
		EXPECT().
		Fetch(&fakePaginationRequestParams).
		Return(nil, fmt.Errorf("ANY ERROR"))

	sut := New(mockProductRepository)
	result, err := sut.Fetch(&fakePaginationRequestParams)

	require.Error(t, err)
	require.Nil(t, result)
}
