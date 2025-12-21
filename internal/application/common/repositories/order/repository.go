package order

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/domain/order"
)

type Repository interface {
	GetByUUID(uuid uuid.UUID) (order.Order, error)
	UpdateOrder(oldOrder order.Order, newOrder order.Order) error
	CreateOrder(order order.Order) error
	RemoveOrder(order order.Order) error
}
