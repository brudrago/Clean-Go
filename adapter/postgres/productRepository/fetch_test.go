package productRepository

import (
	"fmt"
	"testing"

	"github.com/brudrago/clean-go/core/domain"
	"github.com/brudrago/clean-go/core/dto"
	"github.com/bxcodec/faker/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupFetch() ([]string, dto.PaginationRequestParams, domain.Product, pgxmock.PgxPoolIface) {
	cols := []string{"id", "name", "price", "description"}
	fakePaginationRequestParams := dto.PaginationRequestParams{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}
	fakeProductDBResponse := domain.Product{}
	faker.FakeData(&fakeProductDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakePaginationRequestParams, fakeProductDBResponse, mock
}

func TestFetch(t *testing.T) {
	cols, fakePaginationRequestParams, fakeProductDBResponse, mock := setupFetch()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM product").
		WillReturnRows(pgxmock.NewRows(cols).AddRow(
			fakeProductDBResponse.ID,
			fakeProductDBResponse.Name,
			fakeProductDBResponse.Price,
			fakeProductDBResponse.Description,
		))

	mock.ExpectQuery("SELECT COUNT(.+) FROM product").
		WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(int32(1)))

	sut := New(mock)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for _, product := range products.Items {
		require.NotEmpty(t, product.ID)
		require.Equal(t, fakeProductDBResponse.Name, product.Name)
		require.Equal(t, fakeProductDBResponse.Price, product.Price)
		require.Equal(t, fakeProductDBResponse.Description, product.Description)
	}
}

func TestFetch_QueryError(t *testing.T) {
	_, fakePaginationRequestParams, _, mock := setupFetch()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM product").
		WillReturnError(fmt.Errorf("ANY QUERY ERROR"))

	sut := New(mock)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, products)
}

func TestFetch_QueryCountError(t *testing.T) {
	cols, fakePaginationRequestParams, fakeProductDBResponse, mock := setupFetch()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM product").
		WillReturnRows(pgxmock.NewRows(cols).AddRow(
			fakeProductDBResponse.ID,
			fakeProductDBResponse.Name,
			fakeProductDBResponse.Price,
			fakeProductDBResponse.Description,
		))

	mock.ExpectQuery("SELECT COUNT(.+) FROM product").
		WillReturnError(fmt.Errorf("ANY QUERY COUNT ERROR"))

	sut := New(mock)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	require.Nil(t, products)
}
