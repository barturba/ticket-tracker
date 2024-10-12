package data

import (
	"database/sql"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

// Incidents

type Incident struct {
	ID                  uuid.UUID          `json:"id"`
	CreatedAt           time.Time          `json:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at"`
	ShortDescription    string             `json:"short_description"`
	Description         sql.NullString     `json:"description"`
	ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
	CompanyID           uuid.UUID          `json:"company_id"`
	AssignedToID        uuid.NullUUID      `json:"assigned_to_id"`
	State               database.StateEnum `json:"state"`
}

func ValidateIncident(v *validator.Validator, incident *Incident) {
	v.Check(incident.ID != uuid.UUID{}, "id", "must be provided")

	v.Check(incident.ShortDescription != "", "short_description", "must be provided")
	v.Check(len(incident.ShortDescription) <= 500, "short_description", "must not be more than 500 bytes long")

	v.Check(len(incident.Description.String) <= 50000, "description", "must not be more than 50 kBytes long")

	v.Check(incident.ConfigurationItemID != uuid.UUID{}, "configuration_item_id", "must be provided")

	v.Check(incident.CompanyID != uuid.UUID{}, "company_id", "must be provided")

	v.Check(incident.State != "", "state", "must be provided")

	var stateStrings []string

	// Iterate over enum values and convert to strings
	for _, s := range []database.StateEnum{
		database.StateEnumNew,
		database.StateEnumInProgress,
		database.StateEnumAssigned,
		database.StateEnumOnHold,
		database.StateEnumResolved,
	} {
		stateStrings = append(stateStrings, string(s))
	}
	v.Check(validator.PermittedValue(string(incident.State), stateStrings...), "state", "invalid state value")

}
