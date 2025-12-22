package manageProduct

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type UnmarkProductAsOutOfStockInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *UnmarkProductAsOutOfStockInteractor) Execute(uuid uuid.UUID) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return ErrProductNotFound
	}

	err = product.UnmarkAsOutOfStock()
	if err != nil {
		return err
	}

	err = i.Repository.Update(product)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.ProductUnmarkOutOfStockEvent{
		UUID: product.UUID,
	})
	if err != nil {
		return err
	}

	return nil
}
