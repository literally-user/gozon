package manageOrder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type MarkCompletedInteractor struct {
	Repository repositories.OrderRepository
}

func (i *MarkCompletedInteractor) Execute(orderUUID uuid.UUID) error {
	order, err := i.Repository.GetByUUID(orderUUID)
	if err != nil {
		return fmt.Errorf("mark completed interactor: failed to get order by uuid: %w", err)
	}

	err = order.MarkAsCompleted()
	if err != nil {
		return fmt.Errorf("mark completed interactor: failed to mark order as completed: %w", err)
	}

	err = i.Repository.Update(order)
	if err != nil {
		return fmt.Errorf("mark completed interactor: failed to update order: %w", err)
	}

	return nil
}
