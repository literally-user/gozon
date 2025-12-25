package manageUser

import (
	"crypto/sha256"
	"fmt"

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
		return fmt.Errorf("change password: failed to get user by uuid: %w", err)
	}

	oldPassword := user.Password
	newHashedPassword := sha256.Sum256([]byte(password))

	if err = user.ChangePassword(password); err != nil {
		return fmt.Errorf("change password: failed to change password: %w", err)
	}

	err = i.Publisher.Publish(publisher.UserChangedPasswordEvent{
		UUID:        uuid,
		OldPassword: oldPassword,
		NewPassword: newHashedPassword,
	})
	if err != nil {
		return fmt.Errorf("change password: failed to publish: %w", err)
	}

	return nil
}
