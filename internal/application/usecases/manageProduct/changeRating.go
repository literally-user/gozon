package manageProduct

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeProductRatingInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *ChangeProductRatingInteractor) Execute(uuid uuid.UUID, rating float32) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return fmt.Errorf("change product rating: failed to get product by uuid: %w", err)
	}

	oldRating := product.Rating()

	err = product.ChangeRating(rating)
	if err != nil {
		return fmt.Errorf("change product rating: failed to change rating: %w", err)
	}

	err = i.Repository.Update(product)
	if err != nil {
		return fmt.Errorf("change product rating: failed to update: %w", err)
	}

	err = i.Publisher.Publish(publisher.ProductChangedRatingEvent{
		UUID:      product.UUID,
		OldRating: oldRating,
		NewRating: rating,
	})
	if err != nil {
		return fmt.Errorf("change product rating: failed to publish: %w", err)
	}

	return nil
}
