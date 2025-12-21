package user

import (
	"crypto/sha256"

	"github.com/google/uuid"
)

type User struct {
	UUID uuid.UUID

	Phone     string
	Username  string
	Password  [32]byte
	Email     string
	Privilege int

	Banned bool
}

func NewUser(username, password, email, phone string) (User, error) {
	var err error

	user := User{
		UUID:      uuid.New(),
		Privilege: 0,
		Banned:    false,
	}

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

	err = user.ChangePhone(phone)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *User) ChangeUsername(username string) error {
	if username == u.Username {
		return ErrUsernameDoesntChanged
	}

	u.Username = username
	return nil
}

func (u *User) ChangePhone(phone string) error {
	if phone == u.Phone {
		return ErrPhoneDoesntChanged
	}

	u.Phone = phone
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

func (u *User) BanUser() error {
	if u.Banned == true {
		return ErrBanStatusDoesntChanged
	}
	u.Banned = true
	return nil
}

func (u *User) UnbanUser() error {
	if u.Banned == false {
		return ErrUnbanStatusDoesntChanged
	}
	u.Banned = false
	return nil
}

func (u *User) ChangePrivilege(privilege int) error {
	if privilege == u.Privilege {
		return ErrPrivilegeDoesntChanged
	}

	u.Privilege = privilege
	return nil
}
