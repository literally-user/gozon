package order

import "github.com/google/uuid"

type Order struct {
	UUID      uuid.UUID
	Cancelled bool
	Completed bool
	Taken     bool

	ProductUUID uuid.UUID
	UserUUID    uuid.UUID
}

func NewOrder(productUuid, userUuid uuid.UUID) (Order, error) {
	delivery := Order{
		UUID: uuid.New(),

		ProductUUID: productUuid,
		UserUUID:    userUuid,
	}

	return delivery, nil
}

func (d *Order) MarkAsCompleted() error {
	if d.Completed == true {
		return ErrCompletedStatusDoesntChanged
	}
	d.Completed = true
	return nil
}

func (d *Order) UnmarkAsCompleted() error {
	if d.Completed == false {
		return ErrCompletedStatusDoesntChanged
	}
	d.Completed = false
	return nil
}

func (d *Order) MarkAsCancelled() error {
	if d.Cancelled == true {
		return ErrCancelledStatusDoesntChanged
	}
	d.Cancelled = true
	return nil
}

func (d *Order) UnmarkAsCancelled() error {
	if d.Cancelled == false {
		return ErrCancelledStatusDoesntChanged
	}
	d.Cancelled = false
	return nil
}

func (d *Order) MarkAsTaken() error {
	if d.Taken == true {
		return ErrTakenStatusDoesntChanged
	}
	d.Taken = true
	return nil
}

func (d *Order) UnmarkAsTaken() error {
	if d.Taken == false {
		return ErrTakenStatusDoesntChanged
	}
	d.Taken = false
	return nil
}
