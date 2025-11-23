package productRepository

import (
	"github.com/brudrago/clean-go/adapter/postgres"
	"github.com/brudrago/clean-go/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
