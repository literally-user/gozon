package manageOrder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeAddressInteractor struct {
	Repository repositories.OrderRepository
}

func (i *ChangeAddressInteractor) Execute(orderUUID uuid.UUID, address string) error {
	order, err := i.Repository.GetByUUID(orderUUID)
	if err != nil {
		return fmt.Errorf("change address interactor: failed to get order by uuid: %w", err)
	}

	err = order.ChangeAddress(address)
	if err != nil {
		return fmt.Errorf("change address interactor: failed to change address: %w", err)
	}

	err = i.Repository.Update(order)
	if err != nil {
		return fmt.Errorf("change address interactor: failed to update order: %w", err)
	}

	return nil
}
