package manageOrder

import (
	"fmt"

	"github.com/literally_user/gozon/internal/application/common/bank"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
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
		return domainOrder.Order{}, fmt.Errorf("create order: failed to get user by uuid: %w", err)
	}
	cartItem, err := i.CartItemRepository.GetByUUID(orderDTO.CartItemUUID)
	if err != nil {
		return domainOrder.Order{}, fmt.Errorf("create order: failed to get cart item by uuid: %w", err)
	}
	product, err := i.ProductRepository.GetByUUID(cartItem.ProductUUID)
	if err != nil {
		return domainOrder.Order{}, fmt.Errorf("create order: failed to get product by uuid: %w", err)
	}

	err = i.CartItemRepository.Remove(cartItem)
	if err != nil {
		return domainOrder.Order{}, fmt.Errorf("create order: failed to remove cart item: %w", err)
	}

	err = product.ChangeCount(product.Count() - 1)
	if err != nil {
		return domainOrder.Order{}, fmt.Errorf("create order: failed to decrease count of product: %w", err)
	}

	err = i.ProductRepository.Update(product)
	if err != nil {
		return domainOrder.Order{}, fmt.Errorf("create order: failed to update product: %w", err)
	}

	bankAdapter, err := i.BankAdapterFactory.GetBankAdapter(orderDTO.BankName)
	if err != nil {
		return domainOrder.Order{}, fmt.Errorf("create order: failed to find adapter by name: %w", err)
	}

	err = bankAdapter.Refund(orderDTO.Card)
	if err != nil {
		return domainOrder.Order{}, fmt.Errorf("create order: failed to refund money: %w", err)
	}

	order, err := domainOrder.NewOrder(orderDTO.Address, product.UUID, user.UUID)
	if err != nil {
		return domainOrder.Order{}, fmt.Errorf("create order: failed to create order: %w", err)
	}

	err = i.Publisher.Publish(publisher.OrderCreatedEvent{
		Order:       order,
		UserUUID:    user.UUID,
		ProductUUID: product.UUID,
	})
	if err != nil {
		return domainOrder.Order{}, fmt.Errorf("create order: failed to publish: %w", err)
	}

	return order, nil
}
