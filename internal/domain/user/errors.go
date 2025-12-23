package user

import "errors"

var ErrUsernameDoesntChanged = errors.New("username doesn't changed")
var ErrUsernameWrongFormat = errors.New("username string is wrong")

var ErrPasswordDoesntChanged = errors.New("password doesn't changed")
var ErrPasswordWrongFormat = errors.New("password string is wrong")

var ErrEmailDoesntChanged = errors.New("email doesn't changed")
var ErrEmailWrongFormat = errors.New("email string is wrong")

var ErrBanStateDoesntChanged = errors.New("ban state doesn't changed")
var ErrPrivilegeDoesntChanged = errors.New("privilege doesn't changed")

var ErrTelephoneDoesntChanged = errors.New("telephone doesn't changed")
var ErrTelephoneWrongFormat = errors.New("telephone number string is wrong")
