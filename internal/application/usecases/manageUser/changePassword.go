package manageUser

import (
	"crypto/sha256"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangePasswordInteractor struct {
	Repository repositories.UserRepository
	Publisher  publisher.Publisher
}

func (i *ChangePasswordInteractor) Execute(uuid uuid.UUID, password string) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return ErrUserNotFound
	}

	oldPassword := user.Password
	newHashedPassword := sha256.Sum256([]byte(password))

	err = user.ChangePassword(password)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.UserChangedPasswordEvent{
		UUID:        uuid,
		OldPassword: oldPassword,
		NewPassword: newHashedPassword,
	})

	return nil
}
