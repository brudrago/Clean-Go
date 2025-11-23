package productService

import (
	"github.com/brudrago/clean-go/core/domain"
)

type service struct {
	useCase domain.ProductUseCase
}

func New(useCase domain.ProductUseCase) domain.ProductService {
	return &service{
		useCase: useCase,
	}
}
