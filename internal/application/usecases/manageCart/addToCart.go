package manageCart

import (
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
	cartItemDomain "github.com/literally_user/gozon/internal/domain/cartItem"
)

type AddToCartInteractor struct {
	CartItemRepository repositories.CartItemRepository
	ProductRepository  repositories.ProductRepository
	Publisher          publisher.Publisher
}

func (i *AddToCartInteractor) Execute(cartItemDTO DTO) (cartItemDomain.CartItem, error) {
	newCartItem, err := cartItemDomain.NewCartItem(cartItemDTO.UserUUID, cartItemDTO.ProductUUID)
	if err != nil {
		return cartItemDomain.CartItem{}, err
	}

	err = i.CartItemRepository.Create(newCartItem)
	if err != nil {
		return cartItemDomain.CartItem{}, err
	}

	product, err := i.ProductRepository.GetByUUID(cartItemDTO.ProductUUID)
	if err != nil {
		return cartItemDomain.CartItem{}, applicationErrors.ErrProductNotFound
	}

	err = product.ChangeShadowRating(product.ProductShadowRating() + 0.1)
	if err != nil {
		return cartItemDomain.CartItem{}, err
	}

	err = i.ProductRepository.Update(product)
	if err != nil {
		return cartItemDomain.CartItem{}, err
	}

	err = i.Publisher.Publish(publisher.CartItemCreatedEvent{
		CartItem: newCartItem,
	})
	if err != nil {
		return cartItemDomain.CartItem{}, err
	}

	return newCartItem, nil
}
