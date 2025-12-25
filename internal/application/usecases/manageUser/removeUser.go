package manageUser

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type DeleteUserInteractor struct {
	Repository repositories.UserRepository
	Publisher  publisher.Publisher
}

func (i *DeleteUserInteractor) Execute(uuid uuid.UUID) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return fmt.Errorf("remove user: failed to get user by uuid: %w", err)
	}

	err = i.Repository.Remove(user)
	if err != nil {
		return fmt.Errorf("remove user: failed to remove user: %w", err)
	}

	err = i.Publisher.Publish(publisher.UserRemovedEvent{
		User: user,
	})
	if err != nil {
		return fmt.Errorf("remove user: failed to publish: %w", err)
	}

	return nil
}
