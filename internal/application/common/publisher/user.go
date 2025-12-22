package publisher

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/domain/user"
)

type UserCreatedEvent struct {
	User user.User
}

func (UserCreatedEvent) Name() string {
	return "create.user"
}

type UserRemovedEvent struct {
	User user.User
}

func (UserRemovedEvent) Name() string {
	return "remove.user"
}

type UserChangedUsernameEvent struct {
	UUID        uuid.UUID
	OldUsername string
	NewUsername string
}

func (UserChangedUsernameEvent) Name() string {
	return "change.user.username"
}

type UserChangedPasswordEvent struct {
	UUID        uuid.UUID
	OldPassword [32]byte
	NewPassword [32]byte
}

func (UserChangedPasswordEvent) Name() string {
	return "change.user.password"
}

type UserChangedEmailEvent struct {
	UUID     uuid.UUID
	OldEmail string
	NewEmail string
}

func (UserChangedEmailEvent) Name() string {
	return "change.user.email"
}

type UserChangedTelephoneEvent struct {
	UUID         uuid.UUID
	OldTelephone string
	NewTelephone string
}

func (UserChangedTelephoneEvent) Name() string {
	return "change.user.telephone"
}
