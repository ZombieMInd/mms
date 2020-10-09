package model_test

import (
	"testing"

	"github.com/ZombieMInd/mms/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestingUser(t)
	assert.NoError(t, u.BeforCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestingUser(t)
			},
			isValid: true,
		},
		{
			name: "encrypted password",
			u: func() *model.User {
				u := model.TestingUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpassword"

				return u
			},
			isValid: true,
		},
		{
			name: "Empty email",
			u: func() *model.User {
				u := model.TestingUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			name: "Invalid email",
			u: func() *model.User {
				u := model.TestingUser(t)
				u.Email = "123"

				return u
			},
			isValid: false,
		},
		{
			name: "Too short password",
			u: func() *model.User {
				u := model.TestingUser(t)
				u.Password = "123"

				return u
			},
			isValid: false,
		},
		{
			name: "Empty password",
			u: func() *model.User {
				u := model.TestingUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
