package manageProduct

import (
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
		return ErrProductNotFound
	}

	oldRating := product.Rating

	err = product.ChangeRating(rating)
	if err != nil {
		return err
	}

	err = i.Repository.Update(product)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.ProductChangedRatingEvent{
		UUID:      product.UUID,
		OldRating: oldRating,
		NewRating: rating,
	})
	if err != nil {
		return err
	}

	return nil
}
