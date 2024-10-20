package data

import (
	"time"

	"github.com/barturba/ticket-tracker/pkg/validator"
	"github.com/google/uuid"
)

// User represents a user in the system with fields for ID, creation and update timestamps, first name, last name, and email.
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role"`
}

// ValidateUser validates the fields of a User struct using the provided validator.
// It checks that:
// - ID is provided and is not an empty UUID.
// - FirstName is not more than 50 bytes long.
// - LastName is not more than 50 bytes long.
// - Email is provided and is not more than 320 bytes long.
func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.ID != uuid.UUID{}, "id", "must be provided")

	v.Check(len(user.FirstName) <= 50, "first_name", "must not be more than 50 bytes long")

	v.Check(len(user.LastName) <= 50, "last_name", "must not be more than 50 bytes long")

	v.Check(user.Email != "", "email", "must be provided")
	v.Check(len(user.Email) <= 320, "email", "must not be more than 320 bytes long")

}
