package order

import "errors"

var ErrCompletedStateDoesntChanged = errors.New("order completed state doesn't changed")
var ErrTakenStateDoesntChanged = errors.New("order taken state doesn't changed")
var ErrAddressDoesntChanged = errors.New("order address destination doesn't changed")
