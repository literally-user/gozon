package comment

import "errors"

var ErrContentDoesntChanged = errors.New("comment content doesn't changed")
var ErrWrongContentFormat = errors.New("comment content wrong format")

var ErrRateDoesntChanged = errors.New("comment rate doesn't changed")
