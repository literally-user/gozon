package repositories

import (
	"github.com/google/uuid"
	domain "github.com/literally_user/gozon/internal/domain/order"
)

type OrderRepository interface {
	GetAllOrders() []domain.Order
	GetByUUID(uuid uuid.UUID) (domain.Order, error)
	Create(user domain.Order) error
	Update(user domain.Order) error
	Remove(user domain.Order) error
}
