package product

import (
	"github.com/literally_user/gozon/internal/application/common/repositories"
	domain "github.com/literally_user/gozon/internal/domain/product"
)

type CreateProductInteractor struct {
	Repository repositories.ProductRepository
}

func (i *CreateProductInteractor) Execute(title, description, productType string, price float32) error {
	product, err := domain.NewProduct(title, description, productType, price)
	if err != nil {
		return err
	}

	err = i.Repository.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}
