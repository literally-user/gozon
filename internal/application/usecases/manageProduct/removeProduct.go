package manageProduct

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type RemoveProductInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *RemoveProductInteractor) Execute(uuid uuid.UUID) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return applicationErrors.ErrProductNotFound
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
