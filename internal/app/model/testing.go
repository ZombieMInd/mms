package model

import "testing"

// TestingUser ...
func TestingUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.com",
		Password: "password",
	}
}
