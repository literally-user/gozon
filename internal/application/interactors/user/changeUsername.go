package user

import (
	"github.com/google/uuid"
	application "github.com/literally_user/gozon/internal/application/common/repositories/user"
)

type ChangeUsernameInteractor struct {
	Repository application.Repository
}

func (i *ChangeUsernameInteractor) Execute(uuid uuid.UUID, username string) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return ErrUserNotFound
	}

	err = user.ChangeUsername(username)
	if err != nil {
		return err
	}

	err = i.Repository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
