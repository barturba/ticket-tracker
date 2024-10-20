// Data provides the data structures and validation logic.
package data

import (
	"time"

	"github.com/barturba/ticket-tracker/internal/utils/validator"
	"github.com/google/uuid"
)

// CI represents a Configuration Item in the ticket-tracker system.
type CI struct {
	ID        uuid.UUID `json:"id"`         // Unique identifier for the CI.
	CreatedAt time.Time `json:"created_at"` // Timestamp when the CI was created.
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the CI was last updated.
	Name      string    `json:"name"`       // Name of the CI.
}

// ValidateCI checks the validity of a CI instance, ensuring that
// the ID is a not empty UUID and that the name is not more than 50 bytes.
func ValidateCI(v *validator.Validator, ci *CI) {
	v.Check(ci.ID != uuid.UUID{}, "id", "must be provided")

	v.Check(len(ci.Name) <= 50, "name", "must not be more than 50 bytes long")
}
