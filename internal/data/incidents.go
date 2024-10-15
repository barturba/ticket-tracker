package data

import (
	"database/sql"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

// Incident represents an incident record in the ticket-tracker system.
// It contains information about the incident such as its ID, creation and update timestamps,
// short and detailed descriptions, associated configuration item and company IDs, assigned user,
// and the current state of the incident.
type Incident struct {
	ID                  uuid.UUID          `json:"id"`
	CreatedAt           time.Time          `json:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at"`
	ShortDescription    string             `json:"short_description"`
	Description         sql.NullString     `json:"description"`
	ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
	CompanyID           uuid.UUID          `json:"company_id"`
	AssignedToID        uuid.NullUUID      `json:"assigned_to_id"`
	AssignedToName      string             `json:"assigned_to_name"`
	State               database.StateEnum `json:"state"`
}

// The ValidateIncident function validates the fields of an Incident instance using the provided
// validator. It ensures that required fields are provided and that string fields do not exceed
// their maximum lengths. It also checks that the state field has a valid value from the predefined
// set of state enums.
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
