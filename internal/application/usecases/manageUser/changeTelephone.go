package manageUser

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type ChangeTelephoneInteractor struct {
	Repository repositories.UserRepository
	Publisher  publisher.Publisher
}

func (i *ChangeTelephoneInteractor) Execute(uuid uuid.UUID, telephone string) error {
	user, err := i.Repository.GetByUUID(uuid)
	if err != nil {
		return fmt.Errorf("change telephone: failed to get user by uuid: %w", err)
	}

	oldTelephone := user.Telephone

	err = user.ChangeTelephone(telephone)
	if err != nil {
		return fmt.Errorf("change telephone: failed to change telephone: %w", err)
	}

	err = i.Publisher.Publish(publisher.UserChangedTelephoneEvent{
		UUID:         uuid,
		OldTelephone: oldTelephone,
		NewTelephone: telephone,
	})
	if err != nil {
		return fmt.Errorf("change telephone: failed to publish: %w", err)
	}

	return nil
}
