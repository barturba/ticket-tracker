package data

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func (f Filters) SortColumn() string {
	for _, safeValue := range f.SortSafelist {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + f.Sort)
}

func (f Filters) SortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
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

// Companies
type Company struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type CompanyRow struct {
	Count     int64     `json:"count"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func ValidateCompany(v *validator.Validator, company *Company) {
	v.Check(company.ID != uuid.UUID{}, "id", "must be provided")

	v.Check(company.Name != "", "name", "must be provided")
	v.Check(len(company.Name) <= 500, "name", "must not be more than 500 bytes long")
}

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

// CIs
type CI struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
}

func ValidateCI(v *validator.Validator, ci *CI) {
	v.Check(ci.ID != uuid.UUID{}, "id", "must be provided")

	v.Check(len(ci.Name) <= 50, "name", "must not be more than 50 bytes long")
}

// Config

type Config struct {
	Host     string
	Port     string
	Env      string
	PageSize int
}

// Define a struct to keep track of page metadata
type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

// CalculateMetadata() calculates the pagination metadata values given the total
// number of records, current page, and page size values.
func CalculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		// Return an empty Metadata struct if there are no records
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     (totalRecords + pageSize - 1) / pageSize,
		TotalRecords: totalRecords,
	}
}
