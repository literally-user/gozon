package manageUser

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeEmailInteractor struct {
	Repository repositories.UserRepository
	Publisher  publisher.Publisher
}

func (i *ChangeEmailInteractor) Execute(uuid uuid.UUID, email string) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return fmt.Errorf("change email: failed to get user by uuid: %w", err)
	}

	oldEmail := user.Email

	err = user.ChangeEmail(email)
	if err != nil {
		return fmt.Errorf("change email: failed to change email: %w", err)
	}

	err = i.Publisher.Publish(publisher.UserChangedEmailEvent{
		UUID:     uuid,
		OldEmail: oldEmail,
		NewEmail: email,
	})
	if err != nil {
		return fmt.Errorf("change email: failed to publish: %w", err)
	}

	return nil
}
