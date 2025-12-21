package order

import "errors"

var ErrCompletedStatusDoesntChanged error = errors.New("delivered status doesn't changed")
var ErrCancelledStatusDoesntChanged error = errors.New("cancelled status doesn't changed")
var ErrTakenStatusDoesntChanged error = errors.New("taken status doesn't changed")
