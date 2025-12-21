package events

import (
	"github.com/literally_user/gozon/internal/domain/product"
	"github.com/literally_user/gozon/internal/domain/user"
)

type CreatedOrderEvent struct {
	User    user.User
	Product product.Product
}

func (CreatedOrderEvent) Name() string {
	return "order.created"
}
