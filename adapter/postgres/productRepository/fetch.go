package productRepository

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/brudrago/clean-go/core/domain"
	"github.com/brudrago/clean-go/core/dto"
)

func (repo repository) Fetch(pagination *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	ctx := context.Background()
	products := []domain.Product{}
	var total int32

	// Builder base
	builder := paginate.NewBuilder().
		Table("product p").
		Model(&domain.Product{}).
		Page(int(pagination.Page)).
		Limit(int(pagination.ItemsPerPage))

	// Busca (search)
	if pagination.Search != "" {
		builder = builder.Search(pagination.Search, "name", "description")
	}

	// Sort (primeiro campo da lista)
	if len(pagination.Sort) > 0 {
		sortField := pagination.Sort[0]

		// Se houver qualquer coisa em Descending, vamos considerar que é descendente
		isDescending := len(pagination.Descending) > 0

		if isDescending {
			builder = builder.OrderByDesc(sortField)
		} else {
			builder = builder.OrderBy(sortField)
		}
	}

	// SQL da listagem
	query, args, err := builder.BuildSQL()
	if err != nil {
		return nil, err
	}

	// SQL do COUNT
	countQuery, countArgs, err := builder.BuildCountSQL()
	if err != nil {
		return nil, err
	}

	// Busca paginada
	rows, err := repo.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := domain.Product{}
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Description,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// COUNT total
	if err := repo.db.QueryRow(ctx, countQuery, countArgs...).Scan(&total); err != nil {
		return nil, err
	}

	paginationResult := &domain.Pagination[[]domain.Product]{
		Items: products,
		Total: total,
	}

	return paginationResult, nil
}
