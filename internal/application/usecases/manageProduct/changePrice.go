package manageProduct

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeProductPriceInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *ChangeProductPriceInteractor) Execute(uuid uuid.UUID, price float64) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return fmt.Errorf("change product description: failed to get product by uuid: %w", err)
	}

	oldPrice := product.Price()

	err = product.ChangePrice(price)
	if err != nil {
		return fmt.Errorf("change product price: failed to change price: %w", err)
	}

	err = i.Repository.Update(product)
	if err != nil {
		return fmt.Errorf("change product price: failed to update: %w", err)
	}

	err = i.Publisher.Publish(publisher.ProductChangedPriceEvent{
		UUID:     product.UUID,
		OldPrice: oldPrice,
		NewPrice: price,
	})
	if err != nil {
		return fmt.Errorf("change product price: failed to publish: %w", err)
	}

	return nil
}
