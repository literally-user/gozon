package user

import (
	"crypto/sha256"

	"github.com/google/uuid"
)

type User struct {
	UUID uuid.UUID

	Password  [32]byte
	Username  string
	Email     string
	Telephone string

	Privileges int
	Banned     bool
}

func NewUser(username, password, email, telephone string) (User, error) {
	var err error
	var user User

	err = user.ChangeUsername(username)
	if err != nil {
		return User{}, err
	}

	err = user.ChangePassword(password)
	if err != nil {
		return User{}, err
	}

	err = user.ChangeEmail(email)
	if err != nil {
		return User{}, err
	}

	err = user.ChangeTelephone(telephone)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *User) ChangeTelephone(telephone string) error {
	if telephone == u.Telephone {
		return ErrTelephoneDoesntChanged
	}

	u.Telephone = telephone
	return nil
}

func (u *User) ChangeUsername(username string) error {
	if username == u.Username {
		return ErrUsernameDoesntChanged
	}

	u.Username = username
	return nil
}

func (u *User) ChangePassword(password string) error {
	hashedPassword := sha256.Sum256([]byte(password))

	if hashedPassword == u.Password {
		return ErrPasswordDoesntChanged
	}

	u.Password = hashedPassword
	return nil
}

func (u *User) ChangeEmail(email string) error {
	if email == u.Email {
		return ErrEmailDoesntChanged
	}

	u.Email = email
	return nil
}

func (u *User) Ban() error {
	if u.Banned {
		return ErrBanStateDoesntChanged
	}

	u.Banned = true
	return nil
}

func (u *User) Unban() error {
	if !u.Banned {
		return ErrBanStateDoesntChanged
	}

	u.Banned = false
	return nil
}

func (u *User) SetUserPrivileges() error {
	if u.Privileges == 0 {
		return ErrPrivilegeDoesntChanged
	}

	u.Privileges = 0
	return nil
}

func (u *User) SetAdminPrivileges() error {
	if u.Privileges == 1 {
		return ErrPrivilegeDoesntChanged
	}

	u.Privileges = 1
	return nil
}
