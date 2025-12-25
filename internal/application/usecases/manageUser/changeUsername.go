package manageUser

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeUsernameInteractor struct {
	Repository repositories.UserRepository
	Publisher  publisher.Publisher
}

func (i *ChangeUsernameInteractor) Execute(uuid uuid.UUID, username string) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return fmt.Errorf("change username: failed to get user by uuid: %w", err)
	}

	oldUsername := user.Username

	err = user.ChangeUsername(username)
	if err != nil {
		return fmt.Errorf("change username: failed to change username: %w", err)
	}

	err = i.Publisher.Publish(publisher.UserChangedUsernameEvent{
		UUID:        uuid,
		OldUsername: oldUsername,
		NewUsername: username,
	})
	if err != nil {
		return fmt.Errorf("change username: failed to publish: %w", err)
	}

	return nil
}
