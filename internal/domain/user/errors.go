package user

import "errors"

var ErrUsernameDoesntChanged = errors.New("username doesn't changed")
var ErrPasswordDoesntChanged = errors.New("password doesn't changed")
var ErrEmailDoesntChanged = errors.New("email doesn't changed")
var ErrBanStateDoesntChanged = errors.New("ban state doesn't changed")
var ErrPrivilegeDoesntChanged = errors.New("privilege doesn't changed")
var ErrTelephoneDoesntChanged = errors.New("telephone doesn't changed")
