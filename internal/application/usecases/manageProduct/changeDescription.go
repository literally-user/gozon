package manageProduct

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeProductDescriptionInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *ChangeProductDescriptionInteractor) Execute(uuid uuid.UUID, description string) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return fmt.Errorf("change product description: failed to get product by uuid: %w", err)
	}

	oldDescription := product.Description()

	err = product.ChangeDescription(description)
	if err != nil {
		return fmt.Errorf("change product description: failed to change description: %w", err)
	}

	err = i.Repository.Update(product)
	if err != nil {
		return fmt.Errorf("change product description: failed to update product: %w", err)
	}

	err = i.Publisher.Publish(publisher.ProductChangedDescriptionEvent{
		UUID:           product.UUID,
		OldDescription: oldDescription,
		NewDescription: description,
	})
	if err != nil {
		return fmt.Errorf("change product description: failed to publish: %w", err)
	}

	return nil
}
