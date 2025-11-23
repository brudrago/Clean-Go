package productRepository

import (
	"context"

	"github.com/brudrago/clean-go/core/domain"
	"github.com/brudrago/clean-go/core/dto"
)

func (repo repository) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	ctx := context.Background()
	product := &domain.Product{}

	err := repo.db.QueryRow(
		ctx,
		"INSERT INTO product (name, price, description) VALUES ($1, $2, $3) returning *",
		productRequest.Name,
		productRequest.Price,
		productRequest.Description,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Description,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
