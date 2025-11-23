package productRepository

import (
	"context"

	"github.com/brudrago/clean-go/core/domain"
	"github.com/brudrago/clean-go/core/dto"
	"github.com/brudrago/go-paginate/paginate"
)

func (repo repository) Fetch(pagination *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	ctx := context.Background()
	products := []domain.Product{}
	total := int32(0)

	query, queryCount, err := paginate.Paginate("SELECT * FROM product").
		Page(pagination.Page).
		ItemsPerPage(pagination.ItemsPerPage).
		Search(pagination.Search, []string{"name", "description"}).
		Sort(pagination.Sort).
		Descending(pagination.Descending).
		Query()
	if err != nil {
		return nil, err
	}

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Description,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	err = repo.db.QueryRow(ctx, queryCount).Scan(&total)
	if err != nil {
		return nil, err
	}

	paginationResult := &domain.Pagination[[]domain.Product]{
		Items: products,
		Total: total,
	}

	return paginationResult, nil
}
