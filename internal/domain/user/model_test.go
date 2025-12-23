package user

import (
	"crypto/sha256"
	"testing"
)

func TestUser_ChangeTelephone(t *testing.T) {
	tests := []struct {
		name          string
		initial       string
		newTelephone  string
		wantErr       error
		wantTelephone string
	}{
		{
			name:          "successful telephone change",
			initial:       "+70000000000",
			newTelephone:  "+12345678900",
			wantErr:       nil,
			wantTelephone: "+12345678900",
		},
		{
			name:          "telephone doesn't change - same value",
			initial:       "+70000000000",
			newTelephone:  "+70000000000",
			wantErr:       ErrTelephoneDoesntChanged,
			wantTelephone: "+70000000000",
		},
		{
			name:          "invalid telephone format",
			initial:       "+70000000000",
			newTelephone:  "0000000",
			wantErr:       ErrTelephoneWrongFormat,
			wantTelephone: "+70000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{Telephone: tt.initial}

			err := u.ChangeTelephone(tt.newTelephone)

			if err != tt.wantErr {
				t.Errorf("ChangeTelephone() error = %v, wantErr %v", err, tt.wantErr)
			}

			if u.Telephone != tt.wantTelephone {
				t.Errorf("ChangeTelephone() telephone = %v, want %v", u.Telephone, tt.wantTelephone)
			}
		})
	}
}

func TestUser_ChangeUsername(t *testing.T) {
	tests := []struct {
		name         string
		initial      string
		newUsername  string
		wantErr      error
		wantUsername string
	}{
		{
			name:         "successful username change with underscore",
			initial:      "literally",
			newUsername:  "danger_literally",
			wantErr:      nil,
			wantUsername: "danger_literally",
		},
		{
			name:         "successful username change lowercase",
			initial:      "literally",
			newUsername:  "dangerliterally",
			wantErr:      nil,
			wantUsername: "dangerliterally",
		},
		{
			name:         "successful username change camelCase",
			initial:      "literally",
			newUsername:  "dangerLiterally",
			wantErr:      nil,
			wantUsername: "dangerliterally",
		},
		{
			name:         "successful username change uppercase",
			initial:      "literally",
			newUsername:  "DANGERLITERALLY",
			wantErr:      nil,
			wantUsername: "dangerliterally",
		},
		{
			name:         "username doesn't change - same value",
			initial:      "literally",
			newUsername:  "literally",
			wantErr:      ErrUsernameDoesntChanged,
			wantUsername: "literally",
		},
		{
			name:         "invalid username with special characters",
			initial:      "literally",
			newUsername:  "danger))literally",
			wantErr:      ErrUsernameWrongFormat,
			wantUsername: "literally",
		},
		{
			name:         "invalid username with space",
			initial:      "literally",
			newUsername:  "danger literally",
			wantErr:      ErrUsernameWrongFormat,
			wantUsername: "literally",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{Username: tt.initial}

			err := u.ChangeUsername(tt.newUsername)

			if err != tt.wantErr {
				t.Errorf("ChangeUsername() error = %v, wantErr %v", err, tt.wantErr)
			}

			if u.Username != tt.wantUsername {
				t.Errorf("ChangeUsername() username = %v, want %v", u.Username, tt.wantUsername)
			}
		})
	}
}

func TestUser_ChangePassword(t *testing.T) {
	tests := []struct {
		name         string
		initial      string
		newPassword  string
		wantErr      error
		wantPassword string
	}{
		{
			name:         "successful password change",
			initial:      "Password123!",
			newPassword:  "Password1234!",
			wantErr:      nil,
			wantPassword: "Password1234!",
		},
		{
			name:         "successful password change with multiple special chars",
			initial:      "Password123!",
			newPassword:  "Password12345--+!",
			wantErr:      nil,
			wantPassword: "Password12345--+!",
		},
		{
			name:         "password doesn't change - same value",
			initial:      "Password123!",
			newPassword:  "Password123!",
			wantErr:      ErrPasswordDoesntChanged,
			wantPassword: "Password123!",
		},
		{
			name:         "invalid password - all lowercase",
			initial:      "Password123!",
			newPassword:  "password12345--+!",
			wantErr:      ErrPasswordWrongFormat,
			wantPassword: "Password123!",
		},
		{
			name:         "invalid password - no special characters",
			initial:      "Password123!",
			newPassword:  "Password12345",
			wantErr:      ErrPasswordWrongFormat,
			wantPassword: "Password123!",
		},
		{
			name:         "invalid password - no numbers",
			initial:      "Password123!",
			newPassword:  "Password!!!!",
			wantErr:      ErrPasswordWrongFormat,
			wantPassword: "Password123!",
		},
		{
			name:         "invalid password - too short",
			initial:      "Password123!",
			newPassword:  "Pa1!",
			wantErr:      ErrPasswordWrongFormat,
			wantPassword: "Password123!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{Password: sha256.Sum256([]byte(tt.initial))}

			err := u.ChangePassword(tt.newPassword)

			if err != tt.wantErr {
				t.Errorf("ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
			}

			var wantHashPassword [32]byte
			if tt.wantErr == nil {
				wantHashPassword = sha256.Sum256([]byte(tt.newPassword))
			} else {
				wantHashPassword = sha256.Sum256([]byte(tt.initial))
			}

			if u.Password != wantHashPassword {
				t.Errorf("ChangePassword() password hash mismatch, got = %v, want %v", u.Password, wantHashPassword)
			}
		})
	}
}
func TestUser_ChangeEmail(t *testing.T) {
	tests := []struct {
		name      string
		initial   string
		newEmail  string
		wantErr   error
		wantEmail string
	}{
		{
			name:      "successful email change",
			initial:   "literally@email.io",
			newEmail:  "literally@ltu.io",
			wantErr:   nil,
			wantEmail: "literally@ltu.io",
		},
		{
			name:      "successful email change with long domain",
			initial:   "literally@email.io",
			newEmail:  "llllttttuuuu@llllttttuuuu.iiiooo",
			wantErr:   nil,
			wantEmail: "llllttttuuuu@llllttttuuuu.iiiooo",
		},
		{
			name:      "email doesn't change - same value",
			initial:   "literally@email.io",
			newEmail:  "literally@email.io",
			wantErr:   ErrEmailDoesntChanged,
			wantEmail: "literally@email.io",
		},
		{
			name:      "invalid email format",
			initial:   "literally@email.io",
			newEmail:  "invalid email format",
			wantErr:   ErrEmailWrongFormat,
			wantEmail: "literally@email.io",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{Email: tt.initial}

			err := u.ChangeEmail(tt.newEmail)

			if err != tt.wantErr {
				t.Errorf("ChangeEmail() error = %v, wantErr %v", err, tt.wantErr)
			}

			if u.Email != tt.wantEmail {
				t.Errorf("ChangeEmail() email = %v, want %v", u.Email, tt.wantEmail)
			}
		})
	}
}
