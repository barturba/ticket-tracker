// Data provides the data structures and validation logic.
package data

import (
	"time"

	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

// CI represents a configuration item.
type CI struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

// ValidateCI checks the validity of a CI instance, ensuring that
// the ID is a not empty UUID and that the name is not more than 50 bytes.
func ValidateCI(v *validator.Validator, ci *CI) {
	v.Check(ci.ID != uuid.UUID{}, "id", "must be provided")

	v.Check(len(ci.Name) <= 50, "name", "must not be more than 50 bytes long")
}
