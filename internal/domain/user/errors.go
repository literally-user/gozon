package user

import "errors"

var ErrUsernameDoesntChanged = errors.New("user username doesn't changed")
var ErrUsernameWrongFormat = errors.New("user username string is wrong")

var ErrPasswordDoesntChanged = errors.New("user password doesn't changed")
var ErrPasswordWrongFormat = errors.New("user password string is wrong")

var ErrEmailDoesntChanged = errors.New("user email doesn't changed")
var ErrEmailWrongFormat = errors.New("user email string is wrong")

var ErrBanStateDoesntChanged = errors.New("user ban state doesn't changed")
var ErrPrivilegeDoesntChanged = errors.New("user privilege doesn't changed")

var ErrTelephoneDoesntChanged = errors.New("user telephone doesn't changed")
var ErrTelephoneWrongFormat = errors.New("user telephone number string is wrong")
