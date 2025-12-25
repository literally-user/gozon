package manageCart

import (
	"fmt"

	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
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
		return cartItemDomain.CartItem{}, fmt.Errorf("add to cart: failed to create new domain cart item: %w", err)
	}

	err = i.CartItemRepository.Create(newCartItem)
	if err != nil {
		return cartItemDomain.CartItem{}, fmt.Errorf("add to cart: failed to create new cart item: %w", err)
	}

	product, err := i.ProductRepository.GetByUUID(cartItemDTO.ProductUUID)
	if err != nil {
		return cartItemDomain.CartItem{}, fmt.Errorf("add to cart: failed to find product by uuid: %w", err)
	}

	err = product.ChangeShadowRating(product.ShadowRating() + 0.1)
	if err != nil {
		return cartItemDomain.CartItem{}, fmt.Errorf("add to cart: failed to change shadow rating: %w", err)
	}

	err = i.ProductRepository.Update(product)
	if err != nil {
		return cartItemDomain.CartItem{}, fmt.Errorf("add to cart: failed to update product: %w", err)
	}

	err = i.Publisher.Publish(publisher.CartItemCreatedEvent{
		CartItem: newCartItem,
	})
	if err != nil {
		return cartItemDomain.CartItem{}, fmt.Errorf("add to cart: failed to publish: %w", err)
	}

	return newCartItem, nil
}
