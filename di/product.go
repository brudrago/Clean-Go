package di

import (
	productService "github.com/brudrago/clean-go/adapter/http/productservice"
	"github.com/brudrago/clean-go/adapter/postgres"
	"github.com/brudrago/clean-go/adapter/postgres/productRepository"
	"github.com/brudrago/clean-go/core/domain"
	productUseCase "github.com/brudrago/clean-go/core/domain/usecase/productusecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productRepository.New(conn)
	productUseCase := productUseCase.New(productRepository)
	productService := productService.New(productUseCase)

	return productService
}
