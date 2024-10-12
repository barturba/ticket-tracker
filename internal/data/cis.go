package data

import (
	"time"

	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

// CIs

type CI struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func ValidateCI(v *validator.Validator, ci *CI) {
	v.Check(ci.ID != uuid.UUID{}, "id", "must be provided")

	v.Check(len(ci.Name) <= 50, "name", "must not be more than 50 bytes long")
}
