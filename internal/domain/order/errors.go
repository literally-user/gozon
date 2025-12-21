package order

import "errors"

var ErrCompletedStatusDoesntChanged error = errors.New("delivered status doesn't changed")
var ErrPaidStatusDoesntChanged error = errors.New("payed status doesn't changed")
var ErrTakenStatusDoesntChanged error = errors.New("taken status doesn't changed")
