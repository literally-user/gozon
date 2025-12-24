package manageProduct

import (
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	domain "github.com/literally_user/gozon/internal/domain/product"
)

type GetAllProductsInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *GetAllProductsInteractor) Execute() ([]domain.Product, error) {
	products := i.Repository.GetAllProducts()

	return products, nil
}
