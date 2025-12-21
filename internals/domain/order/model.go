package order

import "github.com/google/uuid"

type Order struct {
	UUID      uuid.UUID
	Completed bool
	Taken     bool

	ProductUUID uuid.UUID
	UserUUID    uuid.UUID
}

func NewOrder(productUuid, userUuid uuid.UUID, completed, taken bool) (Order, error) {
	var err error

	delivery := Order{
		UUID:        uuid.New(),
		ProductUUID: productUuid,
		UserUUID:    userUuid,
	}

	if completed {
		err = delivery.MarkAsCompleted()
		if err != nil {
			return Order{}, err
		}
	}

	if taken {
		err = delivery.MarkAsTaken()
		if err != nil {
			return Order{}, err
		}
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
