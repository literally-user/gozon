package manageProduct

import (
	"fmt"

	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	"github.com/literally_user/gozon/internal/domain/product"
)

type CreateProductInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *CreateProductInteractor) Execute(productDTO DTO) (product.Product, error) {
	newProduct, err := product.NewProduct(productDTO.Title, productDTO.Description, productDTO.ProductType, productDTO.Price)
	if err != nil {
		return product.Product{}, fmt.Errorf("create product: failed to create domain product: %w", err)
	}

	err = i.Repository.Create(newProduct)
	if err != nil {
		return product.Product{}, fmt.Errorf("create product: failed to create product: %w", err)
	}

	err = i.Publisher.Publish(publisher.ProductCreatedEvent{
		Product: newProduct,
	})
	if err != nil {
		return product.Product{}, fmt.Errorf("create product: failed to publish: %w", err)
	}

	return newProduct, nil
}
