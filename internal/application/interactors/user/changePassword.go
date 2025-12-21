package user

import (
	"github.com/google/uuid"
	application "github.com/literally_user/gozon/internal/application/common/repositories/user"
)

type ChangePasswordInteractor struct {
	Repository application.Repository
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
