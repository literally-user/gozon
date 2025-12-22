package manageCart

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	"github.com/literally_user/gozon/internal/domain/cartItem"
)

type AddToCartInteractor struct {
	Repository repositories.CartItemRepository
	Publisher  publisher.Publisher
}

func (i *AddToCartInteractor) Execute(userUUID, product uuid.UUID) error {
	newCartItem, err := cartItem.NewCartItem(userUUID, product)
	if err != nil {
		return err
	}

	err = i.Repository.Create(newCartItem)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.CartItemCreatedEvent{
		CartItem: newCartItem,
	})
	if err != nil {
		return err
	}

	return nil
}
