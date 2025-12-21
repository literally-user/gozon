package order

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type MarkAsCancelledOrderInteractor struct {
	Repository repositories.OrderRepository
}

func (i *MarkAsCancelledOrderInteractor) Execute(orderUUID uuid.UUID) error {
	order, err := i.Repository.GetByUUID(orderUUID)
	if err != nil {
		return ErrOrderNotFound
	}

	err = order.MarkAsCancelled()
	if err != nil {
		return err
	}

	err = i.Repository.UpdateOrder(order)
	if err != nil {
		return err
	}

	return nil
}
