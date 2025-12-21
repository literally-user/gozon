package product

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangePriceProductInteractor struct {
	Repository repositories.ProductRepository
}

func (i *ChangePriceProductInteractor) Execute(productUUID uuid.UUID, price float32) error {
	product, err := i.Repository.GetByUUID(productUUID)
	if err != nil {
		return ErrProductNotFound
	}

	err = product.ChangePrice(price)
	if err != nil {
		return err
	}

	err = i.Repository.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}
