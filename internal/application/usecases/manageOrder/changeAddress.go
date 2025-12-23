package manageOrder

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type ChangeAddressInteractor struct {
	Repository repositories.OrderRepository
}

func (i *ChangeAddressInteractor) Execute(orderUUID uuid.UUID, address string) error {
	order, err := i.Repository.GetByUUID(orderUUID)
	if err != nil {
		return applicationErrors.ErrOrderNotFound
	}

	err = order.ChangeAddress(address)
	if err != nil {
		return err
	}

	err = i.Repository.Update(order)
	if err != nil {
		return err
	}

	return nil
}
