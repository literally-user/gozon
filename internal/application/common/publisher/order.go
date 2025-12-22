package publisher

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/domain/order"
)

type OrderCreatedEvent struct {
	Order       order.Order
	UserUUID    uuid.UUID
	ProductUUID uuid.UUID
}

func (OrderCreatedEvent) Name() string {
	return "create.order"
}

type OrderCanceledEvent struct {
	Order       order.Order
	UserUUID    uuid.UUID
	ProductUUID uuid.UUID
}

func (OrderCanceledEvent) Name() string {
	return "cancel.order"
}
