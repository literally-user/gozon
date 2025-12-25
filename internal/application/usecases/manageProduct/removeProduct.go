package manageProduct

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type RemoveProductInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *RemoveProductInteractor) Execute(uuid uuid.UUID) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return fmt.Errorf("remove product: failed to get product by uuid: %w", err)
	}

	err = i.Repository.Remove(product)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.ProductRemovedEvent{
		Product: product,
	})
	if err != nil {
		return err
	}

	return nil
}
