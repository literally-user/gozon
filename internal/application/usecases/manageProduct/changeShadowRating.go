package manageProduct

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeProductShadowRatingInteractor struct {
	Repository repositories.ProductRepository
	Publisher  publisher.Publisher
}

func (i *ChangeProductShadowRatingInteractor) Execute(uuid uuid.UUID, shadowRating float32) error {
	product, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return ErrProductNotFound
	}

	oldShadowRating := product.ShadowRating

	err = product.ChangeShadowRating(shadowRating)
	if err != nil {
		return err
	}

	err = i.Repository.Update(product)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.ProductChangedShadowRatingEvent{
		UUID:            product.UUID,
		OldShadowRating: oldShadowRating,
		NewShadowRating: shadowRating,
	})
	if err != nil {
		return err
	}

	return nil
}
