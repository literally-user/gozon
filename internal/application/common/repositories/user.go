package repositories

import (
	"github.com/google/uuid"
	domain "github.com/literally_user/gozon/internal/domain/user"
)

type UserRepository interface {
	GetAllUsers() []domain.User
	GetByUUID(uuid uuid.UUID) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	Create(user domain.User) error
	Update(user domain.User) error
	Remove(user domain.User) error
}
