package errors

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrWrongPassword = errors.New("password is wrong")
