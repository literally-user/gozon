package manageProduct

import (
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	"github.com/literally_user/gozon/internal/domain/product"
)

type CreateProductInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *CreateProductInteractor) Execute(productDTO DTO) error {
	newProduct, err := product.NewProduct(productDTO.title, productDTO.description, productDTO.productType, productDTO.price)
	if err != nil {
		return err
	}

	err = i.Repository.Create(newProduct)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.ProductCreatedEvent{
		Product: newProduct,
	})
	if err != nil {
		return err
	}

	return nil
}
