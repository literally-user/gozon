package order

import "github.com/google/uuid"

type Order struct {
	UUID uuid.UUID

	Address string

	Completed bool
	Canceled  bool
	Taken     bool

	ProductUUID uuid.UUID
	UserUUID    uuid.UUID
}

func NewOrder(address string, productUUID, userUUID uuid.UUID) (Order, error) {
	order := Order{
		UUID: uuid.New(),

		UserUUID:    userUUID,
		ProductUUID: productUUID,
	}

	err := order.ChangeAddress(address)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

func (o *Order) ChangeAddress(address string) error {
	if o.Address == address {
		return ErrAddressDoesntChanged
	}

	o.Address = address
	return nil
}

func (o *Order) MarkAsCompleted() error {
	if o.Completed {
		return ErrCompletedStateDoesntChanged
	}

	o.Completed = true
	return nil
}

func (o *Order) MarkAsCanceled() error {
	if o.Canceled {
		return ErrCompletedStateDoesntChanged
	}

	o.Canceled = true
	return nil
}

func (o *Order) MarkAsTaken() error {
	if o.Taken {
		return ErrTakenStateDoesntChanged
	}

	o.Taken = true
	return nil
}
