package user

import (
	"testing"

	"github.com/google/uuid"
)

func TestUser_ChangeTelephone(t *testing.T) {
	type fields struct {
		UUID       uuid.UUID
		Password   [32]byte
		Username   string
		Email      string
		Telephone  string
		Privileges int
		Banned     bool
	}
	type args struct {
		telephone string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				UUID:       tt.fields.UUID,
				Password:   tt.fields.Password,
				Username:   tt.fields.Username,
				Email:      tt.fields.Email,
				Telephone:  tt.fields.Telephone,
				Privileges: tt.fields.Privileges,
				Banned:     tt.fields.Banned,
			}
			if err := u.ChangeTelephone(tt.args.telephone); (err != nil) != tt.wantErr {
				t.Errorf("ChangeTelephone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}