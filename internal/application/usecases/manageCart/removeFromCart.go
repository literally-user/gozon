package manageCart

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type RemoveFromCartInteractor struct {
	Repository repositories.CartItemRepository
	Publisher  publisher.Publisher
}

func (i *RemoveFromCartInteractor) Execute(uuid uuid.UUID) error {
	cartItem, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return err
	}

	err = i.Repository.Remove(cartItem)
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
