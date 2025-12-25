package manageOrder

import (
	"fmt"

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
		return fmt.Errorf("cancel order: failed to get order by uuid: %w", applicationErrors.ErrOrderNotFound)
	}

	user, err := i.UserRepository.GetByUUID(order.UserUUID)
	if err != nil {
		return fmt.Errorf("cancel order: failed to get user by uuid: %w", applicationErrors.ErrUserNotFound)
	}
	product, err := i.ProductRepository.GetByUUID(order.ProductUUID)
	if err != nil {
		return fmt.Errorf("cancel order: failed to get product by uuid: %w", applicationErrors.ErrProductNotFound)
	}

	err = product.ChangeCount(product.Count() + 1)
	if err != nil {
		return fmt.Errorf("cancel order: failed to change count: %w", err)
	}

	err = i.ProductRepository.Update(product)
	if err != nil {
		return fmt.Errorf("cancel order: failed to update product: %w", err)
	}

	bankAdapter, err := i.BankAdapterFactory.GetBankAdapter(bankName)
	if err != nil {
		return fmt.Errorf("cancel order: failed to get bank adapter: %w", applicationErrors.ErrBankNotFound)
	}

	err = bankAdapter.Refund(card)
	if err != nil {
		return fmt.Errorf("cancel order: failed to refund money: %w", err)
	}

	err = order.MarkAsCanceled()
	if err != nil {
		return fmt.Errorf("cancel order: failed to mark as canceled: %w", err)
	}

	err = i.OrderRepository.Update(order)
	if err != nil {
		return fmt.Errorf("cancel order: failed to update order: %w", err)
	}

	err = i.Publisher.Publish(publisher.OrderCanceledEvent{
		Order:       order,
		UserUUID:    user.UUID,
		ProductUUID: product.UUID,
	})
	if err != nil {
		return fmt.Errorf("cancel order: failed to publish: %w", err)
	}

	return nil
}
