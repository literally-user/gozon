package manageUser

import (
	"fmt"

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
		return user.User{}, fmt.Errorf("create user: failed to create domain user: %w", err)
	}

	err = i.Repository.Create(newUser)
	if err != nil {
		return user.User{}, fmt.Errorf("create user: failed to create user: %w", err)
	}

	err = i.Publisher.Publish(publisher.UserCreatedEvent{
		User: newUser,
	})
	if err != nil {
		return user.User{}, fmt.Errorf("create user: failed to publish: %w", err)
	}

	return newUser, nil
}
