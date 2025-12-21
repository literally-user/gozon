package product

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeTitleProductInteractor struct {
	Repository repositories.ProductRepository
}

func (i *ChangeTitleProductInteractor) Execute(productUUID uuid.UUID, title string) error {
	product, err := i.Repository.GetByUUID(productUUID)
	if err != nil {
		return ErrProductNotFound
	}

	err = product.ChangeTitle(title)
	if err != nil {
		return err
	}

	err = i.Repository.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}
