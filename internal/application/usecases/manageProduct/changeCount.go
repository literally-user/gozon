package manageProduct

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type ChangeProductCountInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *ChangeProductCountInteractor) Execute(uuid uuid.UUID, count int) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return applicationErrors.ErrProductNotFound
	}

	oldCount := product.ProductCount()

	err = product.ChangeCount(count)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.ProductChangedCountEvent{
		UUID:     uuid,
		OldCount: oldCount,
		NewCount: count,
	})
	if err != nil {
		return err
	}

	return nil
}
