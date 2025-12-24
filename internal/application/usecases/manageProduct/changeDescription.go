package manageProduct

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type ChangeProductDescriptionInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *ChangeProductDescriptionInteractor) Execute(uuid uuid.UUID, description string) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return applicationErrors.ErrProductNotFound
	}

	oldDescription := product.Description()

	err = product.ChangeDescription(description)
	if err != nil {
		return err
	}

	err = i.Repository.Update(product)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.ProductChangedDescriptionEvent{
		UUID:           product.UUID,
		OldDescription: oldDescription,
		NewDescription: description,
	})
	if err != nil {
		return err
	}

	return nil
}
