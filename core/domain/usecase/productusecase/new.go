package productUseCase

import (
	"github.com/brudrago/clean-go/core/domain"
)

type usecase struct {
	repo domain.ProductRepository
}

func New(repo domain.ProductRepository) domain.ProductUseCase {
	return &usecase{
		repo: repo,
	}
}
