package user

import "errors"

var ErrPasswordDoesntChanged = errors.New("password doesn't changed")
var ErrUsernameDoesntChanged error = errors.New("username doesn't changed")
var ErrEmailDoesntChanged error = errors.New("email doesn't changed")
var ErrBanStatusDoesntChanged error = errors.New("ban status doesn't changed")
var ErrUnbanStatusDoesntChanged error = errors.New("unban status doesn't changed")
var ErrPrivilegeDoesntChanged error = errors.New("privilege doesn't changed")
var ErrPhoneDoesntChanged error = errors.New("phone number doesn't changed")
