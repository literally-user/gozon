package product

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type DeleteProductInteractor struct {
	Repository repositories.ProductRepository
}

func (i *DeleteProductInteractor) Execute(productUUID uuid.UUID) error {
	product, err := i.Repository.GetByUUID(productUUID)
	if err != nil {
		return ErrProductNotFound
	}

	err = i.Repository.RemoveProduct(product)
	if err != nil {
		return err
	}

	return nil
}
