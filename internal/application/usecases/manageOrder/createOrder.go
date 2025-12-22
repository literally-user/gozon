package manageOrder

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/bank"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	cartItemApplication "github.com/literally_user/gozon/internal/application/usecases/manageCart"
	productApplication "github.com/literally_user/gozon/internal/application/usecases/manageProduct"
	userApplication "github.com/literally_user/gozon/internal/application/usecases/manageUser"
	domainOrder "github.com/literally_user/gozon/internal/domain/order"
)

type CreateOrderInteractor struct {
	OrderRepository    repositories.OrderRepository
	UserRepository     repositories.UserRepository
	ProductRepository  repositories.ProductRepository
	CartItemRepository repositories.CartItemRepository

	Publisher publisher.Publisher

	BankAdapterFactory bank.AdapterFactory
}

func (i *CreateOrderInteractor) Execute(userUUID, cartItemUUID uuid.UUID, address string, card bank.Card, bankName string) error {
	user, err := i.UserRepository.GetByUUID(userUUID)
	if err != nil {
		return userApplication.ErrUserNotFound
	}
	cartItem, err := i.CartItemRepository.GetByUUID(cartItemUUID)
	if err != nil {
		return cartItemApplication.ErrCartItemNotFound
	}
	product, err := i.ProductRepository.GetByUUID(cartItem.ProductUUID)
	if err != nil {
		return productApplication.ErrProductNotFound
	}

	err = i.CartItemRepository.Remove(cartItem)
	if err != nil {
		return err
	}

	err = product.ChangeCount(product.Count - 1)
	if err != nil {
		return err
	}

	err = i.ProductRepository.Update(product)
	if err != nil {
		return err
	}

	bankAdapter, err := i.BankAdapterFactory.GetBankAdapter(bankName)
	if err != nil {
		return err
	}

	err = bankAdapter.Refund(card)
	if err != nil {
		return err
	}

	order, err := domainOrder.NewOrder(address, product.UUID, user.UUID)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.OrderCreatedEvent{
		Order:       order,
		UserUUID:    user.UUID,
		ProductUUID: product.UUID,
	})

	return nil
}
