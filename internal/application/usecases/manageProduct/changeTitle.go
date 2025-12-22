package manageProduct

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeProductTitleInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *ChangeProductTitleInteractor) Execute(uuid uuid.UUID, title string) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return ErrProductNotFound
	}

	oldTitle := product.Title

	err = product.ChangeTitle(title)
	if err != nil {
		return err
	}

	err = i.Repository.Update(product)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.ProductChangedTitleEvent{
		UUID:     product.UUID,
		OldTitle: oldTitle,
		NewTitle: title,
	})
	if err != nil {
		return err
	}

	return nil
}
