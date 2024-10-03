package models

import (
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type ConfigurationItem struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type Incident struct {
	ID                    uuid.UUID          `json:"id"`
	CreatedAt             time.Time          `json:"created_at"`
	UpdatedAt             time.Time          `json:"updated_at"`
	ShortDescription      string             `json:"short_description"`
	Description           string             `json:"description"`
	State                 database.StateEnum `json:"state"`
	AssignedTo            uuid.UUID          `json:"assigned_to_id"`
	AssignedToName        string             `json:"assigned_to_name"`
	ConfigurationItemID   uuid.UUID          `json:"configuration_item_id"`
	ConfigurationItemName string             `json:"configuration_item_name"`
	CompanyID             uuid.UUID          `json:"company_id"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	APIkey    string    `json:"api_key"`
	Token     string    `json:"token"`
}

func DatabaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		APIkey:    user.Apikey,
	}
}

func DatabaseUsersToUsers(users []database.User) []User {
	var items []User
	for _, item := range users {
		items = append(items, DatabaseUserToUser(item))
	}
	return items
}

func DatabaseConfigurationItemToConfigurationItem(configurationItem database.ConfigurationItem) ConfigurationItem {
	return ConfigurationItem{
		ID:        configurationItem.ID,
		CreatedAt: configurationItem.CreatedAt,
		UpdatedAt: configurationItem.UpdatedAt,
		Name:      configurationItem.Name,
	}
}

func DatabaseConfigurationItemsToConfigurationItems(configurationItems []database.ConfigurationItem) []ConfigurationItem {
	var items []ConfigurationItem
	for _, configurationItem := range configurationItems {
		items = append(items, DatabaseConfigurationItemToConfigurationItem(configurationItem))
	}
	return items
}

func DatabaseCompanyToCompany(company database.Company) Company {
	return Company{
		ID:        company.ID,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
		Name:      company.Name,
	}
}
func DatabaseCompaniesToCompanies(companies []database.Company) []Company {
	var items []Company
	for _, item := range companies {
		items = append(items, DatabaseCompanyToCompany(item))
	}
	return items
}

func DatabaseIncidentToIncident(incident database.Incident) Incident {
	return Incident{
		ID:                  incident.ID,
		CreatedAt:           incident.CreatedAt,
		UpdatedAt:           incident.UpdatedAt,
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description.String,
		State:               incident.State,
		AssignedTo:          incident.AssignedTo.UUID,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	}
}

func DatabaseIncidentRowToIncident(incident database.GetIncidentsRow) Incident {
	return Incident{
		ID:                  incident.ID,
		CreatedAt:           incident.CreatedAt,
		UpdatedAt:           incident.UpdatedAt,
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description.String,
		State:               incident.State,
		AssignedTo:          incident.AssignedTo.UUID,
		AssignedToName:      incident.Name.String,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	}
}

func DatabaseIncidentsRowToIncidents(incidents []database.GetIncidentsRow) []Incident {
	var items []Incident
	for _, item := range incidents {
		items = append(items, DatabaseIncidentRowToIncident(item))
	}
	return items
}

func DatabaseIncidentBySearchTermRowToIncident(incident database.GetIncidentsBySearchTermRow) Incident {
	return Incident{
		ID:                  incident.ID,
		CreatedAt:           incident.CreatedAt,
		UpdatedAt:           incident.UpdatedAt,
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description.String,
		State:               incident.State,
		AssignedTo:          incident.AssignedTo.UUID,
		AssignedToName:      incident.Name.String,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	}
}

func DatabaseIncidentsBySearchTermToIncidents(incidents []database.GetIncidentsBySearchTermRow) []Incident {
	var items []Incident
	for _, item := range incidents {
		items = append(items, DatabaseIncidentBySearchTermRowToIncident(item))
	}
	return items
}

func DatabaseIncidentsToIncidents(incidents []database.Incident) []Incident {
	var items []Incident
	for _, item := range incidents {
		items = append(items, DatabaseIncidentToIncident(item))
	}
	return items
}

func DatabaseIncidentBySearchTermLimitOffsetRowToIncident(incident database.GetIncidentsBySearchTermLimitOffsetRow) Incident {
	return Incident{
		ID:                  incident.ID,
		CreatedAt:           incident.CreatedAt,
		UpdatedAt:           incident.UpdatedAt,
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description.String,
		State:               incident.State,
		AssignedTo:          incident.AssignedTo.UUID,
		AssignedToName:      incident.Name.String,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	}
}
func DatabaseIncidentsBySearchTermLimitOffsetRowToIncidents(incidents []database.GetIncidentsBySearchTermLimitOffsetRow) []Incident {
	var items []Incident
	for _, item := range incidents {
		items = append(items, DatabaseIncidentBySearchTermLimitOffsetRowToIncident(item))
	}
	return items
}

type MenuItem struct {
	Name string
	Link string
}

type MenuItems []MenuItem

func NewMenuItem(name, link string) MenuItem {
	return MenuItem{
		Name: name,
		Link: link,
	}
}

type SelectOption struct {
	Name string
	Link string
}

type SelectOptions []SelectOption

func NewSelectOption(name, link string) SelectOption {
	return SelectOption{
		Name: name,
		Link: link,
	}
}

var StateOptionsEnum = []database.StateEnum{
	database.StateEnumNew,
	database.StateEnumAssigned,
	database.StateEnumInProgress,
	database.StateEnumOnHold,
	database.StateEnumResolved,
}

type FormData struct {
	Errors map[string]string
	Values map[string]string
}

func NewFormData() FormData {
	return FormData{
		Errors: map[string]string{},
		Values: map[string]string{},
	}
}

type Page struct {
	Title            string
	Logo             string
	Alert            Alert
	IsLoggedIn       bool
	IsError          bool
	Msg              string
	User             string
	Email            string
	ProfilePicture   string
	MenuItems        MenuItems
	ProfileMenuItems MenuItems
}

// form := NewIncidentForm("PUT", iEPath, selectOptionsCompany, selectOptionsCI, stateOptions, incident, formData)
type IncidentForm struct {
	Action     string
	Path       string
	CancelPath string
	Companies  SelectOptions
	CIs        SelectOptions
	States     SelectOptions
	AssignedTo SelectOptions
	Incident   Incident
	FormData   FormData
}

// "PUT", iEPath, selectOptionsCompany, selectOptionsCI, stateOptions, incident, formData
func NewIncidentForm(action, path, cancelPath string, companies, cis, states, assignedTo SelectOptions, incident Incident, formData FormData) IncidentForm {
	return IncidentForm{
		Action:     action,
		Path:       path,
		CancelPath: cancelPath,
		Companies:  companies,
		CIs:        cis,
		States:     states,
		AssignedTo: assignedTo,
		Incident:   incident,
		FormData:   formData,
	}
}

type Dropdown struct {
	ID            string
	Label         string
	Name          string
	SelectOptions SelectOptions
	Selected      string
	ErrorText     string
	HXGet         string
	HXTarget      string
}

func NewDropdown(id string, label string, selectOptions SelectOptions, selected, errorText, hxGet, hxTarget string) Dropdown {
	return Dropdown{
		ID:            id,
		Label:         label,
		SelectOptions: selectOptions,
		Selected:      selected,
		ErrorText:     errorText,
		HXGet:         hxGet,
		HXTarget:      hxTarget,
	}
}

type InputField struct {
	ID           string
	Label        string
	Value        string
	InputType    string
	Autocomplete string
	ErrorText    string
	HXGet        string
	HXTarget     string
	Disabled     bool
}

func NewInputField(id string, label string, value string, inputType string, autocomplete string, errorText string, hxGet, hxTarget string) InputField {
	return InputField{
		ID:           id,
		Label:        label,
		Value:        value,
		InputType:    inputType,
		Autocomplete: autocomplete,
		ErrorText:    errorText,
		HXGet:        hxGet,
		HXTarget:     hxTarget,
	}
}

type InputFieldDisabled struct {
	ID           string
	Label        string
	Value        string
	InputType    string
	Autocomplete string
	ErrorText    string
	HXGet        string
	HXTarget     string
}

func NewInputFieldDisabled(id string, label string, value string, inputType string, autocomplete string, errorText string, hxGet, hxTarget string) InputFieldDisabled {
	return InputFieldDisabled{
		ID:           id,
		Label:        label,
		Value:        value,
		InputType:    inputType,
		Autocomplete: autocomplete,
		ErrorText:    errorText,
		HXGet:        hxGet,
		HXTarget:     hxTarget,
	}
}

type Field interface {
	Field()
	GetID() string
	GetLabel() string
	SetError(e string)
	GetError() string
}

func (*InputField) Field()         {}
func (*InputFieldDisabled) Field() {}
func (*Dropdown) Field()           {}

func (i *InputField) GetID() string         { return i.ID }
func (i *InputFieldDisabled) GetID() string { return i.ID }
func (i *Dropdown) GetID() string           { return i.ID }

func (i *InputField) GetLabel() string         { return i.Label }
func (i *InputFieldDisabled) GetLabel() string { return i.Label }
func (i *Dropdown) GetLabel() string           { return i.Label }

func (i *InputField) SetError(e string)         { i.ErrorText = e }
func (i *InputFieldDisabled) SetError(e string) { i.ErrorText = e }
func (i *Dropdown) SetError(e string)           { i.ErrorText = e }

func (i *InputField) GetError() string         { return i.ErrorText }
func (i *InputFieldDisabled) GetError() string { return i.ErrorText }
func (i *Dropdown) GetError() string           { return i.ErrorText }

var IncidentInput struct {
	ID                  uuid.UUID          `json:"id"`
	ShortDescription    string             `json:"short_description"`
	Description         string             `json:"description"`
	CompanyID           uuid.UUID          `json:"company_id"`
	AssignedToID        uuid.UUID          `json:"assigned_to_id"`
	ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
	State               database.StateEnum `json:"state"`
}

func CheckIncident(i Incident) map[string]string {
	StateEnum := []database.StateEnum{
		database.StateEnumNew,
		database.StateEnumAssigned,
		database.StateEnumInProgress,
		database.StateEnumOnHold,
		database.StateEnumResolved,
	}
	v := validator.New()
	v.Check(i.ID != uuid.UUID{}, "id", "must be provided")
	v.Check(i.CompanyID != uuid.UUID{}, "company_id", "must be provided")
	v.Check(i.ConfigurationItemID != uuid.UUID{}, "configuration_item_id", "must be provided")
	v.Check(validator.PermittedValue(i.State, StateEnum...), "state", "This field must equal New, Assigned, In Progress, On Hold, or Resolved")
	v.Check(i.AssignedTo != uuid.UUID{}, "assigned_to_id", "must be provided")
	v.Check(i.Description != "", "description", "must be provided")
	v.Check(i.ShortDescription != "", "short_description", "must be provided")
	if !v.Valid() {
		return v.Errors
	}
	return nil
}

type AlertEnum string

const (
	AlertEnumError   AlertEnum = "Error"
	AlertEnumWarning AlertEnum = "Warning"
	AlertEnumSuccess AlertEnum = "Success"
	AlertEnumInfo    AlertEnum = "Info"
)

type Alert struct {
	Message   string
	AlertType AlertEnum
	Color     string
}

func NewAlert(message string, alertType AlertEnum, color string) Alert {
	return Alert{
		Message:   message,
		AlertType: alertType,
		Color:     color,
	}
}
