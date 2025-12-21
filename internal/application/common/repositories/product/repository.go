package product

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/domain/product"
)

type Repository interface {
	GetByUUID(uuid uuid.UUID) (product.Product, error)
	UpdateProduct(oldProduct product.Product, newProduct product.Product) error
	CreateProduct(product product.Product) error
	RemoveProduct(product product.Product) error
}
