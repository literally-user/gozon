package user

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangePasswordInteractor struct {
	Repository repositories.UserRepository
}

func (i *ChangePasswordInteractor) Execute(uuid uuid.UUID, password string) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return ErrUserNotFound
	}

	err = user.ChangePassword(password)
	if err != nil {
		return err
	}

	err = i.Repository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
