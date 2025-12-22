package publisher

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/domain/user"
)

type UserCreatedEvent struct {
	User user.User
}

func (UserCreatedEvent) Name() string {
	return "user.created"
}

type UserRemovedEvent struct {
	User user.User
}

func (UserRemovedEvent) Name() string {
	return "user.removed"
}

type UserChangedUsernameEvent struct {
	UUID        uuid.UUID
	OldUsername string
	NewUsername string
}

func (UserChangedUsernameEvent) Name() string {
	return "user.change.username"
}

type UserChangedPasswordEvent struct {
	UUID        uuid.UUID
	OldPassword [32]byte
	NewPassword [32]byte
}

func (UserChangedPasswordEvent) Name() string {
	return "user.change.password"
}

type UserChangedEmailEvent struct {
	UUID     uuid.UUID
	OldEmail string
	NewEmail string
}

func (UserChangedEmailEvent) Name() string {
	return "user.change.email"
}

type UserChangedTelephoneEvent struct {
	UUID         uuid.UUID
	OldTelephone string
	NewTelephone string
}

func (UserChangedTelephoneEvent) Name() string {
	return "user.change.telephone"
}
