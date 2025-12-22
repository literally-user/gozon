package manageOrder

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type MarkCompletedInteractor struct {
	Repository repositories.OrderRepository
}

func (i *MarkCompletedInteractor) Execute(orderUUID uuid.UUID) error {
	order, err := i.Repository.GetByUUID(orderUUID)
	if err != nil {
		return applicationErrors.ErrOrderNotFound
	}

	err = order.MarkAsCompleted()
	if err != nil {
		return err
	}

	err = i.Repository.Update(order)
	if err != nil {
		return err
	}

	return nil
}
