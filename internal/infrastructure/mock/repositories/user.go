package repositories

import (
	"errors"

	"github.com/google/uuid"
	domain "github.com/literally_user/gozon/internal/domain/user"
)

var storage = make([]domain.User, 0)

type ImMemoryUserRepository struct{}

func NewInMemoryUserRepository() ImMemoryUserRepository {
	return ImMemoryUserRepository{}
}

func (ImMemoryUserRepository) Create(user domain.User) error {
	for _, val := range storage {
		if val.UUID == user.UUID {
			return errors.New("already has")
		}
	}
	return nil
}

func (ImMemoryUserRepository) Update(user domain.User) error {
	for i, val := range storage {
		if val.UUID == user.UUID {
			storage[i] = user
			return nil
		}
	}

	return errors.New("not found")
}

func (ImMemoryUserRepository) Remove(user domain.User) error {
	for i, val := range storage {
		if val.UUID == user.UUID {
			storage[i] = domain.User{}
			return nil
		}
	}
	return errors.New("not found")
}

func (ImMemoryUserRepository) GetByUUID(uuid uuid.UUID) (domain.User, error) {
	for _, val := range storage {
		if val.UUID == uuid {
			return val, nil
		}
	}
	return domain.User{}, errors.New("not found")
}

func (ImMemoryUserRepository) GetAllUsers() []domain.User {
	return storage
}
