package productUseCase

import (
	"github.com/brudrago/clean-go/core/domain"
	"github.com/brudrago/clean-go/core/dto"
)

func (usecase usecase) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product, err := usecase.repo.Create(productRequest)

	if err != nil {
		return nil, err
	}

	return product, nil
}
