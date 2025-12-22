package repositories

import (
	"github.com/google/uuid"
	domain "github.com/literally_user/gozon/internal/domain/product"
)

type ProductRepository interface {
	GetAllProducts() []domain.Product
	GetByUUID(uuid uuid.UUID) (domain.Product, error)
	Create(user domain.Product) error
	Update(user domain.Product) error
	Remove(user domain.Product) error
}
