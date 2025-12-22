package manageUser

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
)

type ChangeTelephoneInteractor struct {
	Repository repositories.UserRepository
	Publisher  publisher.Publisher
}

func (i *ChangeTelephoneInteractor) Execute(uuid uuid.UUID, telephone string) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return applicationErrors.ErrUserNotFound
	}

	oldTelephone := user.Telephone

	err = user.ChangeTelephone(telephone)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.UserChangedTelephoneEvent{
		UUID:         uuid,
		OldTelephone: oldTelephone,
		NewTelephone: telephone,
	})

	return nil
}
