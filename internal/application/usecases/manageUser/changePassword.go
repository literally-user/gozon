package manageUser

import (
	"crypto/sha256"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type ChangePasswordInteractor struct {
	Repository repositories.UserRepository
	Publisher  publisher.Publisher
}

func (i *ChangePasswordInteractor) Execute(uuid uuid.UUID, password string) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return applicationErrors.ErrUserNotFound
	}

	oldPassword := user.Password
	newHashedPassword := sha256.Sum256([]byte(password))

	if err := user.ChangePassword(password); err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.UserChangedPasswordEvent{
		UUID:        uuid,
		OldPassword: oldPassword,
		NewPassword: newHashedPassword,
	})
	if err != nil {
		return err
	}

	return nil
}
