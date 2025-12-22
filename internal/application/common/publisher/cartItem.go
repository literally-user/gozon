package publisher

import (
	"github.com/literally_user/gozon/internal/domain/cartItem"
)

type CartItemCreatedEvent struct {
	CartItem cartItem.CartItem
}

func (CartItemCreatedEvent) Name() string {
	return "create.cart_item"
}

type CartItemRemovedEvent struct {
	CartItem cartItem.CartItem
}

func (CartItemRemovedEvent) Name() string {
	return "remove.cart_item"
}
