package user

import "errors"

var ErrUserNotFound error = errors.New("user not found")
var ErrNotEnoughRights error = errors.New("not enough rights for this operation")
var ErrUserBanned error = errors.New("user is banned")
