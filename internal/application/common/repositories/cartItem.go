package repositories

import (
	"github.com/google/uuid"
	domain "github.com/literally_user/gozon/internal/domain/cartItem"
)

type CartItemRepository interface {
	GetAllByUserUUID(uuid uuid.UUID) []domain.CartItem
	GetByUUID(uuid uuid.UUID) (domain.CartItem, error)
	Create(user domain.CartItem) error
	Update(user domain.CartItem) error
	Remove(user domain.CartItem) error
}
