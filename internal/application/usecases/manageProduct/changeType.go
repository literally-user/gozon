package manageProduct

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeProductTypeInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *ChangeProductTypeInteractor) Execute(uuid uuid.UUID, productType string) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return fmt.Errorf("change product type: failed to get product by uuid: %w", err)
	}

	oldType := product.Type()

	err = product.ChangeType(productType)
	if err != nil {
		return fmt.Errorf("change product type: failed to change type: %w", err)
	}

	err = i.Repository.Update(product)
	if err != nil {
		return fmt.Errorf("change product type: failed to update product: %w", err)
	}

	err = i.Publisher.Publish(publisher.ProductChangedTypeEvent{
		UUID:    product.UUID,
		OldType: oldType,
		NewType: productType,
	})
	if err != nil {
		return fmt.Errorf("change product type: failed to update: %w", err)
	}

	return nil
}
