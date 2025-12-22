package manageUser

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type ChangeUsernameInteractor struct {
	Repository repositories.UserRepository
	Publisher  publisher.Publisher
}

func (i *ChangeUsernameInteractor) Execute(uuid uuid.UUID, username string) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return applicationErrors.ErrUserNotFound
	}

	oldUsername := user.Username

	err = user.ChangeUsername(username)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.UserChangedUsernameEvent{
		UUID:        uuid,
		OldUsername: oldUsername,
		NewUsername: username,
	})

	return nil
}
