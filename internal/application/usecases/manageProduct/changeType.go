package manageProduct

import (
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
		return ErrProductNotFound
	}

	oldType := product.Type

	err = product.ChangeType(productType)
	if err != nil {
		return err
	}

	err = i.Repository.Update(product)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.ProductChangedTypeEvent{
		UUID:    product.UUID,
		OldType: oldType,
		NewType: productType,
	})
	if err != nil {
		return err
	}

	return nil
}
