// models provides the data structures and validation logic.
package models

import (
	"strings"
	"time"
	"unicode/utf8"

	"github.com/barturba/ticket-tracker/internal/utils/validator"
	"github.com/google/uuid"
)

// CI represents a Configuration Item in the ticket-tracker system.
// A Configuration Item is any component that needs to be managed and maintained
// in an IT infrastructure.
type CI struct {
	ID        uuid.UUID `json:"id"`         // Unique identifier for the CI
	CreatedAt time.Time `json:"created_at"` // Timestamp when the CI was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the CI was last updated
	Name      string    `json:"name"`       // Name of the CI (required, 3-50 chars)
}

// ValidateCI checks the validity of a CI instance, ensuring that
// the ID is a not empty UUID and that the name is not more than 50 bytes.
func ValidateCI(v *validator.Validator, ci *CI) {
	// Check required fields
	v.Check(ci.ID != uuid.UUID{}, "id", "must be provided")

	// Validate Name
	v.Check(strings.TrimSpace(ci.Name) != "", "name", "must not be empty")
	v.Check(utf8.RuneCountInString(ci.Name) >= 3, "name", "must be at least 3 characters long")
	v.Check(utf8.RuneCountInString(ci.Name) <= 50, "name", "must not be more than 50 characters long")

	// Validate timestamps
	v.Check(!ci.CreatedAt.IsZero(), "created_at", "must be provided")
	v.Check(!ci.UpdatedAt.IsZero(), "updated_at", "must be provided")

	v.Check(len(ci.Name) <= 50, "name", "must not be more than 50 bytes long")
}
