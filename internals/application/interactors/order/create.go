package order

import (
	"github.com/google/uuid"

	applicationProduct "github.com/literally_user/gozon/internals/application/interactors/product"
	applicationUser "github.com/literally_user/gozon/internals/application/interactors/user"

	repositoryOrder "github.com/literally_user/gozon/internals/application/common/repositories/order"
	repositoryProduct "github.com/literally_user/gozon/internals/application/common/repositories/product"
	repositoryUser "github.com/literally_user/gozon/internals/application/common/repositories/user"
)

type CreateOrderInteractor struct {
	UserRepository    repositoryUser.Repository
	ProductRepository repositoryProduct.Repository
	OrderRepository   repositoryOrder.Repository
}

func (i *CreateOrderInteractor) Execute(userUUID, productUUID uuid.UUID) error {
	user, err := i.UserRepository.GetByUUID(userUUID)
	if err != nil {
		return applicationUser.ErrUserNotFound
	}
	product, err := i.ProductRepository.GetByUUID(productUUID)
	if err != nil {
		return applicationProduct.ErrProductNotFound
	}

	return nil
}
