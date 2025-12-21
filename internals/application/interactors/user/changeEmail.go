package user

import (
	"github.com/google/uuid"
	application "github.com/literally_user/gozon/internals/application/common/repositories/user"
)

type ChangeEmailInteractor struct {
	Repository application.Repository
}

func (i *ChangeEmailInteractor) Execute(uuid uuid.UUID, email string) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return ErrUserNotFound
	}

	err = user.ChangeEmail(email)
	if err != nil {
		return err
	}

	err = i.Repository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
