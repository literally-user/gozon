package manageUser

import (
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	"github.com/literally_user/gozon/internal/domain/user"
)

type CreateUserInteractor struct {
	Repository repositories.UserRepository
	Publisher  publisher.Publisher
}

func (i *CreateUserInteractor) Execute(userDTO DTO) (user.User, error) {
	newUser, err := user.NewUser(userDTO.Username, userDTO.Password, userDTO.Email, userDTO.Telephone)
	if err != nil {
		return user.User{}, err
	}

	err = i.Repository.Create(newUser)
	if err != nil {
		return user.User{}, err
	}

	err = i.Publisher.Publish(publisher.UserCreatedEvent{
		User: newUser,
	})
	if err != nil {
		return user.User{}, err
	}

	return newUser, nil
}
