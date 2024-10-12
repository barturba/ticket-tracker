package data

import (
	"time"

	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

// Users
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	APIkey    string    `json:"api_key"`
	Password  string    `json:"-"`
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.ID != uuid.UUID{}, "id", "must be provided")

	v.Check(len(user.FirstName) <= 50, "first_name", "must not be more than 50 bytes long")

	v.Check(len(user.LastName) <= 50, "last_name", "must not be more than 50 bytes long")

	v.Check(user.APIkey != "", "api_key", "must be provided")
	v.Check(len(user.APIkey) == 64, "api_key", "must not be more than 64 bytes long")

	v.Check(user.Email != "", "email", "must be provided")
	v.Check(len(user.Email) <= 320, "email", "must not be more than 320 bytes long")

	v.Check(len(user.Password) <= 255, "password", "must not be more than 255 bytes long")
}
