package user

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"
	"unicode"

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

	user.UUID = uuid.New()

	err = user.ChangeUsername(username)
	if err != nil {
		return User{}, fmt.Errorf("failed to change username: %w", err)
	}

	err = user.ChangePassword(password)
	if err != nil {
		return User{}, fmt.Errorf("failed to change password: %w", err)
	}

	err = user.ChangeEmail(email)
	if err != nil {
		return User{}, fmt.Errorf("failed to change email: %w", err)
	}

	err = user.ChangeTelephone(telephone)
	if err != nil {
		return User{}, fmt.Errorf("failed to change telephone: %w", err)
	}

	return user, nil
}

func (u *User) ChangeTelephone(telephone string) error {
	match, _ := regexp.Match(`^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$`, []byte(telephone))
	if !match {
		return ErrTelephoneWrongFormat
	}
	if telephone == u.Telephone {
		return ErrTelephoneDoesntChanged
	}

	u.Telephone = telephone
	return nil
}

func (u *User) ChangeUsername(username string) error {
	for _, val := range username {
		if isSpecial(val) {
			return ErrUsernameWrongFormat
		}
	}
	if len(username) < 5 {
		return ErrUsernameWrongFormat
	}
	if username == u.Username {
		return ErrUsernameDoesntChanged
	}

	u.Username = strings.ToLower(username)
	return nil
}

func (u *User) ChangePassword(password string) error {
	var upperChars int
	var specialChars int
	var numChars int

	for _, val := range password {
		if unicode.IsUpper(val) {
			upperChars++
		}
		if unicode.IsPunct(val) {
			specialChars++
		}
		if unicode.IsNumber(val) {
			numChars++
		}
	}
	if upperChars == 0 || specialChars == 0 || numChars == 0 {
		return ErrPasswordWrongFormat
	}
	if len(password) < 5 {
		return ErrPasswordWrongFormat
	}
	hashedPassword := sha256.Sum256([]byte(password))

	if hashedPassword == u.Password {
		return ErrPasswordDoesntChanged
	}

	u.Password = hashedPassword
	return nil
}

func (u *User) ChangeEmail(email string) error {
	match, _ := regexp.Match(`[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+`, []byte(email))
	if !match {
		return ErrEmailWrongFormat
	}
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

func isSpecial(r rune) bool {
	if r != '_' {
		return unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsSpace(r)
	}

	return false
}
