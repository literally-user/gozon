package manageProduct

import (
	"fmt"

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
		return fmt.Errorf("change product title: failed to get product by uuid: %w", err)
	}

	oldTitle := product.Title()

	err = product.ChangeTitle(title)
	if err != nil {
		return fmt.Errorf("change product title: failed to change title: %w", err)
	}

	err = i.Repository.Update(product)
	if err != nil {
		return fmt.Errorf("change product title: failed to update: %w", err)
	}

	err = i.Publisher.Publish(publisher.ProductChangedTitleEvent{
		UUID:     product.UUID,
		OldTitle: oldTitle,
		NewTitle: title,
	})
	if err != nil {
		return fmt.Errorf("change product title: failed to publish: %w", err)
	}

	return nil
}
