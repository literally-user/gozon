package order

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type MarkAsCompletedOrderInteractor struct {
	Repository repositories.OrderRepository
}

func (i *MarkAsCompletedOrderInteractor) Execute(orderUUID uuid.UUID) error {
	order, err := i.Repository.GetByUUID(orderUUID)
	if err != nil {
		return ErrOrderNotFound
	}

	err = order.MarkAsCompleted()
	if err != nil {
		return err
	}

	err = i.Repository.UpdateOrder(order)
	if err != nil {
		return err
	}

	return nil
}
