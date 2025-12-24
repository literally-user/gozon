package repositories

import (
	"errors"

	"github.com/google/uuid"
	userApplication "github.com/literally_user/gozon/internal/application/errors"
	userDomain "github.com/literally_user/gozon/internal/domain/user"
)

func NewUserStorage() []userDomain.User {
	return make([]userDomain.User, 0)
}

type ImMemoryUserRepository struct {
	Storage []userDomain.User
}

func NewInMemoryUserRepository(storage []userDomain.User) ImMemoryUserRepository {
	return ImMemoryUserRepository{
		Storage: storage,
	}
}

func (r *ImMemoryUserRepository) Create(user userDomain.User) error {
	for _, val := range r.Storage {
		if val.UUID == user.UUID {
			return errors.New("already has")
		}
	}
	r.Storage = append(r.Storage, user)
	return nil
}

func (r *ImMemoryUserRepository) Update(user userDomain.User) error {
	for i, val := range r.Storage {
		if val.UUID == user.UUID {
			r.Storage[i] = user
			return nil
		}
	}

	return userApplication.ErrUserNotFound
}

func (r *ImMemoryUserRepository) Remove(user userDomain.User) error {
	for i, val := range r.Storage {
		if val.UUID == user.UUID {
			r.Storage[i] = userDomain.User{}
			return nil
		}
	}
	return userApplication.ErrUserNotFound
}

func (r *ImMemoryUserRepository) GetByUUID(uuid uuid.UUID) (userDomain.User, error) {
	for _, val := range r.Storage {
		if val.UUID == uuid {
			return val, nil
		}
	}
	return userDomain.User{}, userApplication.ErrUserNotFound
}

func (r *ImMemoryUserRepository) GetAllUsers() []userDomain.User {
	return r.Storage
}
