package models

import (
	"testing"
)

func over256String() string {
	tooLongString := ""
	for i := 0; i < 256; i++ {
		tooLongString += "a"
	}
	return tooLongString
}

func TestValidate(t *testing.T) {
	type input struct {
		Name     string
		Email    string
		Password string
	}
	tests := []struct {
		name    string
		fields  input
		wantErr bool
	}{
		{
			name: "valid",
			fields: input{
				Name:     "test",
				Email:    "test@test.com",
				Password: "Test1234",
			},
			wantErr: false,
		},
		{
			name: "not exist name",
			fields: input{
				Name:     "",
				Email:    "test@test.com",
				Password: "Test1234",
			},
			wantErr: true,
		},
		{
			name: "too long name",
			fields: input{
				Name:     over256String(),
				Email:    "test@test.com",
				Password: "Test1234",
			},
			wantErr: true,
		},
		{
			name: "not exist password",
			fields: input{
				Name:     "test",
				Email:    "",
				Password: "Test1234",
			},
			wantErr: true,
		},
		{
			name: "not exist email",
			fields: input{
				Name:     "test",
				Email:    "",
				Password: "Test1234",
			},
			wantErr: true,
		},
		{
			name: "invalid email",
			fields: input{
				Name:     "test",
				Email:    "test",
				Password: "Test1234",
			},
			wantErr: true,
		},
		{
			name: "not exist password",
			fields: input{
				Name:     "test",
				Email:    "test@test.com",
				Password: "",
			},
			wantErr: true,
		},
		{
			name: "too short password",
			fields: input{
				Name:     "test",
				Email:    "test@test.com",
				Password: "Test1",
			},
			wantErr: true,
		},
		{
			name: "too long password",
			fields: input{
				Name:     "test",
				Email:    "test@test.com",
				Password: over256String(),
			},
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			u := User{
				Name:     testCase.fields.Name,
				Email:    testCase.fields.Email,
				Password: testCase.fields.Password,
			}
			if err := u.Validate(); (err != nil) != testCase.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, testCase.wantErr)
			}
		})
	}
}
