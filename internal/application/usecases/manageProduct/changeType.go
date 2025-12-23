package manageProduct

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type ChangeProductTypeInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *ChangeProductTypeInteractor) Execute(uuid uuid.UUID, productType string) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return applicationErrors.ErrProductNotFound
	}

	oldType := product.ProductType()

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
