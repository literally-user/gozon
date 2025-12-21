package user

import (
	"github.com/literally_user/gozon/internal/domain/user"

	"github.com/google/uuid"
)

type Repository interface {
	GetByUUID(uuid uuid.UUID) (user.User, error)
	UpdateUser(user user.User) error
	CreateUser(user user.User) error
	RemoveUser(user user.User) error
}
