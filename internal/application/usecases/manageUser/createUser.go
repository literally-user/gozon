package manageUser

import (
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	applicationErrors "github.com/literally_user/gozon/internal/application/errors"
	"github.com/literally_user/gozon/internal/domain/user"
)

type CreateUserInteractor struct {
	Repository repositories.UserRepository
	Publisher  publisher.Publisher
}

func (i *CreateUserInteractor) Execute(userDTO DTO) error {
	newUser, err := user.NewUser(userDTO.Username, userDTO.Password, userDTO.Email, userDTO.Telephone)
	if err != nil {
		return applicationErrors.ErrUserNotFound
	}

	err = i.Repository.Create(newUser)
	if err != nil {
		return err
	}

	err = i.Publisher.Publish(publisher.UserCreatedEvent{
		User: newUser,
	})
	if err != nil {
		return err
	}

	return nil
}
