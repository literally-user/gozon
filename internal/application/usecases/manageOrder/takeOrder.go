package manageOrder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type TakeOrderInteractor struct {
	Repository repositories.OrderRepository
}

func (i *TakeOrderInteractor) Execute(orderUUID uuid.UUID) error {
	order, err := i.Repository.GetByUUID(orderUUID)
	if err != nil {
		return fmt.Errorf("take order: failed to get order by uuid: %w", err)
	}

	err = order.MarkAsTaken()
	if err != nil {
		return fmt.Errorf("take order: failed to make as taken: %w", err)
	}

	err = i.Repository.Update(order)
	if err != nil {
		return fmt.Errorf("take order: failed to update order: %w", err)
	}

	return nil
}
