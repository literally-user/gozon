package product

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeProductTypeProductInteractor struct {
	Repository repositories.ProductRepository
}

func (i *ChangeProductTypeProductInteractor) Execute(productUUID uuid.UUID, productType string) error {
	product, err := i.Repository.GetByUUID(productUUID)
	if err != nil {
		return ErrProductNotFound
	}

	err = product.ChangeProductType(productType)
	if err != nil {
		return err
	}

	err = i.Repository.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}
