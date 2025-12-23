package manageCart

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type RemoveFromCartInteractor struct {
	CartItemRepository repositories.CartItemRepository
	ProductRepository  repositories.ProductRepository
	Publisher          publisher.Publisher
}

func (i *RemoveFromCartInteractor) Execute(uuid uuid.UUID) error {
	cartItem, err := i.CartItemRepository.GetByUUID(uuid)
	if err != nil {
		return applicationErrors.ErrCartItemNotFound
	}

	err = i.CartItemRepository.Remove(cartItem)
	if err != nil {
		return err
	}

	product, err := i.ProductRepository.GetByUUID(cartItem.ProductUUID)
	if err != nil {
		return applicationErrors.ErrProductNotFound
	}

	err = product.ChangeShadowRating(product.ShadowRating() - 0.1)
	if err != nil {
		return err
	}

	err = i.ProductRepository.Update(product)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.CartItemRemovedEvent{
		CartItem: cartItem,
	})
	if err != nil {
		return err
	}

	return nil
}
