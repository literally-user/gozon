package user

import (
	"github.com/google/uuid"
	application "github.com/literally_user/gozon/internal/application/common/repositories/user"
)

type UnbanUserInteractor struct {
	Repository application.Repository
}

func (i *UnbanUserInteractor) Execute(callerUserUUID uuid.UUID, targetUserUUID uuid.UUID) error {
	callerUser, err := i.Repository.GetByUUID(callerUserUUID)
	if err != nil {
		return ErrUserNotFound
	}
	targetUser, err := i.Repository.GetByUUID(targetUserUUID)
	if err != nil {
		return ErrUserNotFound
	}

	if callerUser.Privilege == 0 {
		return ErrNotEnoughRights
	}

	err = targetUser.UnbanUser()
	if err != nil {
		return err
	}

	err = i.Repository.UpdateUser(targetUser)
	if err != nil {
		return err
	}

	return nil
}
