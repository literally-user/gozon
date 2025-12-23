package user

import (
	"testing"
)

func TestUser_ChangeTelephone(t *testing.T) {
	tests := []struct {
		name      string
		telephone string
		wantErr   bool
	}{
		{
			name:      "Correct telephone",
			telephone: "+12345678900",
			wantErr:   false,
		},
		{
			name:      "Incorrect telephone",
			telephone: "0000000",
			wantErr:   true,
		},
		{
			name:      "Duplicate telephone",
			telephone: "+70000000000",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := NewUser("literally", "Password123!", "literally@email.io", "+70000000000")

			if err := u.ChangeTelephone(tt.telephone); (err != nil) != tt.wantErr {
				t.Errorf("ChangeTelephone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_ChangeUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{
			name:     "Correct username",
			username: "danger_literally",
			wantErr:  false,
		},
		{
			name:     "Correct username",
			username: "dangerliterally",
			wantErr:  false,
		},
		{
			name:     "Correct username",
			username: "dangerLiterally",
			wantErr:  false,
		},
		{
			name:     "Correct username",
			username: "DANGERLITERALLY",
			wantErr:  false,
		},
		{
			name:     "Incorrect username (special characters)",
			username: "danger))literally",
			wantErr:  true,
		},
		{
			name:     "Incorrect username (space)",
			username: "danger literally",
			wantErr:  true,
		},
		{
			name:     "Incorrect username",
			username: "literally",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := NewUser("literally", "Password123!", "literally@email.io", "+70000000000")

			if err := u.ChangeUsername(tt.username); (err != nil) != tt.wantErr {
				t.Errorf("ChangeTelephone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_ChangePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: "Password1234!",
			wantErr:  false,
		},
		{
			name:     "Correct password",
			password: "Password12345--+!",
			wantErr:  false,
		},
		{
			name:     "Incorrect password (all lower)",
			password: "password12345--+!",
			wantErr:  true,
		},
		{
			name:     "Incorrect password (without special characters)",
			password: "Password12345",
			wantErr:  true,
		},
		{
			name:     "Incorrect password (without numbers)",
			password: "Password!!!!",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := NewUser("literally", "Password123!", "literally@email.io", "+70000000000")

			if err := u.ChangePassword(tt.password); (err != nil) != tt.wantErr {
				t.Errorf("ChangeTelephone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
