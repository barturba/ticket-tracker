package data

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/validator" // New import
	"github.com/google/uuid"
)

// Add a SortSafelist field to hold the supported sort values.
type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

func (f Filters) Limit() int {
	return f.PageSize
}

func (f Filters) Offset() int {
	return (f.Page - 1) * f.PageSize
}

func ReadInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		v.AddError(key, "must be an integer value")
		return defaultValue
	}

	return i
}

func ReadString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	return s
}

func ValidateFilters(v *validator.Validator, f Filters) {

	// Check that the page and page_size parameters contain sensible values.
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be no more than ten million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be no more than 100")

	// Check that the sort parameter matches a value in the safelist.
	v.Check(validator.PermittedValue(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}

type Envelope map[string]any

func WriteJSON(w http.ResponseWriter, status int, data Envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// Incidents

type Incident struct {
	ID                  uuid.UUID          `json:"id"`
	CreatedAt           time.Time          `json:"-"`
	UpdatedAt           time.Time          `json:"-"`
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

// Companies
type Company struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
}

func ValidateCompany(v *validator.Validator, company *Company) {
	v.Check(company.ID != uuid.UUID{}, "id", "must be provided")

	v.Check(company.Name != "", "name", "must be provided")
	v.Check(len(company.Name) <= 500, "name", "must not be more than 500 bytes long")
}
