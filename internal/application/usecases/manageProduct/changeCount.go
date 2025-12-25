package manageProduct

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeProductCountInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *ChangeProductCountInteractor) Execute(uuid uuid.UUID, count int) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return fmt.Errorf("change product count: failed to get product by uuid: %w", err)
	}

	oldCount := product.Count()

	err = product.ChangeCount(count)
	if err != nil {
		return fmt.Errorf("change product count: failed to change product count: %w", err)
	}

	err = i.Publisher.Publish(publisher.ProductChangedCountEvent{
		UUID:     uuid,
		OldCount: oldCount,
		NewCount: count,
	})
	if err != nil {
		return fmt.Errorf("change product count: failed to publish: %w", err)
	}

	return nil
}
