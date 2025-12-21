package repositories

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/domain/product"
)

type ProductRepository interface {
	GetByUUID(uuid uuid.UUID) (product.Product, error)
	UpdateProduct(product product.Product) error
	CreateProduct(product product.Product) error
	RemoveProduct(product product.Product) error
}
