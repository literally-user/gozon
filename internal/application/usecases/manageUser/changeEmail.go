package manageUser

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type ChangeEmailInteractor struct {
	Repository repositories.UserRepository
	Publisher  publisher.Publisher
}

func (i *ChangeEmailInteractor) Execute(uuid uuid.UUID, email string) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return applicationErrors.ErrUserNotFound
	}

	oldEmail := user.Email

	err = user.ChangeEmail(email)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.UserChangedEmailEvent{
		UUID:     uuid,
		OldEmail: oldEmail,
		NewEmail: email,
	})

	return nil
}
