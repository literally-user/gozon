package user

import (
	"github.com/literally_user/gozon/internal/application/common/repositories"
	domain "github.com/literally_user/gozon/internal/domain/user"
)

type CreateUserInteractor struct {
	Repository repositories.UserRepository
}

func (i *CreateUserInteractor) Execute(username, password, email, phone string) error {
	user, err := domain.NewUser(username, password, email, phone)
	if err != nil {
		return ErrUserNotFound
	}

	err = i.Repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
