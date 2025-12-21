package order

import (
	"github.com/google/uuid"

	"github.com/literally_user/gozon/internal/application/common/events"
	"github.com/literally_user/gozon/internal/application/common/repositories"

	applicationProduct "github.com/literally_user/gozon/internal/application/interactors/product"
	applicationUser "github.com/literally_user/gozon/internal/application/interactors/user"
	domainOrder "github.com/literally_user/gozon/internal/domain/order"
)

type CreateOrderInteractor struct {
	UserRepository    repositories.UserRepository
	ProductRepository repositories.ProductRepository
	OrderRepository   repositories.OrderRepository
	EventBus          events.EventBus
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

	if user.Banned {
		return applicationUser.ErrUserBanned
	}

	order, err := domainOrder.NewOrder(productUUID, userUUID)
	if err != nil {
		return err
	}

	err = i.OrderRepository.CreateOrder(order)
	if err != nil {
		return err
	}

	err = i.EventBus.Notify(events.CreatedOrderEvent{
		User:    user,
		Product: product,
	})
	if err != nil {
		return err
	}

	return nil
}
