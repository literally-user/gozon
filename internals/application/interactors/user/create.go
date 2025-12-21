package user

import (
	application "github.com/literally_user/gozon/internals/application/common/repositories/user"
	domain "github.com/literally_user/gozon/internals/domain/user"
)

type CreateUserInteractor struct {
	Repository application.Repository
}

func (i *CreateUserInteractor) Execute(username, password, email string) error {
	user, err := domain.NewUser(username, password, email)
	if err != nil {
		return ErrUserNotFound
	}

	err = i.Repository.CreateUser(user)
	if err != nil {
		return ErrUserNotFound
	}

	return nil
}
