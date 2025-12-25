package manageCart

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type RemoveFromCartInteractor struct {
	CartItemRepository repositories.CartItemRepository
	ProductRepository  repositories.ProductRepository
	Publisher          publisher.Publisher
}

func (i *RemoveFromCartInteractor) Execute(uuid uuid.UUID) error {
	cartItem, err := i.CartItemRepository.GetByUUID(uuid)
	if err != nil {
		return fmt.Errorf("remove from cart: failed to get cart item by uuid: %w", err)
	}

	err = i.CartItemRepository.Remove(cartItem)
	if err != nil {
		return fmt.Errorf("remove from cart: failed to remove cart item: %w", err)
	}

	product, err := i.ProductRepository.GetByUUID(cartItem.ProductUUID)
	if err != nil {
		return fmt.Errorf("remove from cart: failed to get product by uuid: %w", err)
	}

	err = product.ChangeShadowRating(product.ShadowRating() - 0.1)
	if err != nil {
		return fmt.Errorf("remove from cart: failed to change shadow rating: %w", err)
	}

	err = i.ProductRepository.Update(product)
	if err != nil {
		return fmt.Errorf("remove from cart: failed to update product: %w", err)
	}

	err = i.Publisher.Publish(publisher.CartItemRemovedEvent{
		CartItem: cartItem,
	})
	if err != nil {
		return fmt.Errorf("remove from cart: failed to publish: %w", err)
	}

	return nil
}
