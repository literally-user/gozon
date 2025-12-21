package repositories

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/domain/order"
)

type OrderRepository interface {
	GetByUUID(uuid uuid.UUID) (order.Order, error)
	UpdateOrder(order order.Order) error
	CreateOrder(order order.Order) error
	RemoveOrder(order order.Order) error
}
