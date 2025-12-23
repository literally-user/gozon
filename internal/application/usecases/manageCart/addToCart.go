package manageCart

import (
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
	"github.com/literally_user/gozon/internal/domain/cartItem"
)

type AddToCartInteractor struct {
	CartItemRepository repositories.CartItemRepository
	ProductRepository  repositories.ProductRepository
	Publisher          publisher.Publisher
}

func (i *AddToCartInteractor) Execute(cartItemDTO DTO) error {
	newCartItem, err := cartItem.NewCartItem(cartItemDTO.UserUUID, cartItemDTO.ProductUUID)
	if err != nil {
		return err
	}

	err = i.CartItemRepository.Create(newCartItem)
	if err != nil {
		return err
	}

	product, err := i.ProductRepository.GetByUUID(cartItemDTO.ProductUUID)
	if err != nil {
		return applicationErrors.ErrProductNotFound
	}

	err = product.ChangeShadowRating(product.ShadowRating() + 0.1)
	if err != nil {
		return err
	}

	err = i.ProductRepository.Update(product)
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
