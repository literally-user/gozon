package manageOrder

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/bank"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type CancelOrderInteractor struct {
	OrderRepository    repositories.OrderRepository
	UserRepository     repositories.UserRepository
	ProductRepository  repositories.ProductRepository
	CartItemRepository repositories.CartItemRepository

	Publisher publisher.Publisher

	BankAdapterFactory bank.AdapterFactory
}

func (i *CancelOrderInteractor) Execute(orderUUID uuid.UUID, bankName string, card bank.Card) error {
	order, err := i.OrderRepository.GetByUUID(orderUUID)
	if err != nil {
		return applicationErrors.ErrOrderNotFound
	}

	user, err := i.UserRepository.GetByUUID(order.UserUUID)
	if err != nil {
		return applicationErrors.ErrUserNotFound
	}
	product, err := i.ProductRepository.GetByUUID(order.ProductUUID)
	if err != nil {
		return applicationErrors.ErrProductNotFound
	}

	err = product.ChangeCount(product.GetCount() + 1)
	if err != nil {
		return err
	}

	err = i.ProductRepository.Update(product)
	if err != nil {
		return err
	}

	bankAdapter, err := i.BankAdapterFactory.GetBankAdapter(bankName)
	if err != nil {
		return applicationErrors.ErrBankNotFound
	}

	err = bankAdapter.Refund(card)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.OrderCanceledEvent{
		Order:       order,
		UserUUID:    user.UUID,
		ProductUUID: product.UUID,
	})

	return nil
}
