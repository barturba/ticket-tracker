package data

import (
	"time"

	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

// Users
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email,omitempty"`
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.ID != uuid.UUID{}, "id", "must be provided")

	v.Check(len(user.FirstName) <= 50, "first_name", "must not be more than 50 bytes long")

	v.Check(len(user.LastName) <= 50, "last_name", "must not be more than 50 bytes long")

	v.Check(user.Email != "", "email", "must be provided")
	v.Check(len(user.Email) <= 320, "email", "must not be more than 320 bytes long")

}
