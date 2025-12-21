package order

import "errors"

var ErrCompletedStateDoesntChanged = errors.New("completed state doesn't changed")
var ErrCanceledStateDoesntChanged = errors.New("canceled state doesn't changed")
var ErrTakenStateDoesntChanged = errors.New("taken state doesn't changed")
var ErrAddressDoesntChanged = errors.New("address destination doesn't changed")
