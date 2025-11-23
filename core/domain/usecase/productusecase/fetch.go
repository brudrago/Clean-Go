package productUseCase

import (
	"github.com/brudrago/clean-go/core/domain"
	"github.com/brudrago/clean-go/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	products, err := usecase.repo.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}
