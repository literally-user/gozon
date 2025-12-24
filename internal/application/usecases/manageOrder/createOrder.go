package manageOrder

import (
	"github.com/literally_user/gozon/internal/application/common/bank"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
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

func (i *CreateOrderInteractor) Execute(orderDTO DTO) (domainOrder.Order, error) {
	user, err := i.UserRepository.GetByUUID(orderDTO.UserUUID)
	if err != nil {
		return domainOrder.Order{}, applicationErrors.ErrUserNotFound
	}
	cartItem, err := i.CartItemRepository.GetByUUID(orderDTO.CartItemUUID)
	if err != nil {
		return domainOrder.Order{}, applicationErrors.ErrCartItemNotFound
	}
	product, err := i.ProductRepository.GetByUUID(cartItem.ProductUUID)
	if err != nil {
		return domainOrder.Order{}, applicationErrors.ErrProductNotFound
	}

	err = i.CartItemRepository.Remove(cartItem)
	if err != nil {
		return domainOrder.Order{}, err
	}

	err = product.ChangeCount(product.ProductCount() - 1)
	if err != nil {
		return domainOrder.Order{}, err
	}

	err = i.ProductRepository.Update(product)
	if err != nil {
		return domainOrder.Order{}, err
	}

	bankAdapter, err := i.BankAdapterFactory.GetBankAdapter(orderDTO.BankName)
	if err != nil {
		return domainOrder.Order{}, applicationErrors.ErrBankNotFound
	}

	err = bankAdapter.Refund(orderDTO.Card)
	if err != nil {
		return domainOrder.Order{}, err
	}

	order, err := domainOrder.NewOrder(orderDTO.Address, product.UUID, user.UUID)
	if err != nil {
		return domainOrder.Order{}, err
	}

	err = i.Publisher.Publish(publisher.OrderCreatedEvent{
		Order:       order,
		UserUUID:    user.UUID,
		ProductUUID: product.UUID,
	})
	if err != nil {
		return domainOrder.Order{}, err
	}

	return order, nil
}
