package user

import (
	"github.com/google/uuid"
	application "github.com/literally_user/gozon/internal/application/common/repositories/user"
)

type DeleteUserInteractor struct {
	Repository application.Repository
}

func (i *DeleteUserInteractor) Execute(callerUserUUID, targetUserUUID uuid.UUID) error {
	callerUser, err := i.Repository.GetByUUID(callerUserUUID)
	if err != nil {
		return ErrUserNotFound
	}
	targetUser, err := i.Repository.GetByUUID(targetUserUUID)
	if err != nil {
		return ErrUserNotFound
	}

	if callerUser.Privilege == 0 {
		if callerUserUUID == targetUserUUID {
			err = i.Repository.RemoveUser(callerUser)
			if err != nil {
				return err
			}
		} else {
			return ErrNotEnoughRights
		}
	}

	err = i.Repository.RemoveUser(targetUser)
	if err != nil {
		return err
	}

	return nil
}
