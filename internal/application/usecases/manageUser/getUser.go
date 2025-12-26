package manageUser

import (
	"crypto/sha256"
	"fmt"

	"github.com/literally_user/gozon/internal/application/common/repositories"
	"github.com/literally_user/gozon/internal/application/errors"
	domain "github.com/literally_user/gozon/internal/domain/user"
)

type GetUserInteractor struct {
	Repository repositories.UserRepository
}

func (i *GetUserInteractor) Execute(username, password string) (domain.User, error) {
	user, err := i.Repository.GetByUsername(username)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: failed to get user by username: %w", err)
	}

	hashedPassword := sha256.Sum256([]byte(password))

	if hashedPassword != user.Password {
		return domain.User{}, fmt.Errorf("get user: failed to authorize: %w", errors.ErrWrongPassword)
	}

	return user, nil
}
