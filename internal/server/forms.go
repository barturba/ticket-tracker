package server

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/models"
	"github.com/barturba/ticket-tracker/views"
)

func MakeIncidentFields(incident models.Incident, companies, cis, states, users models.SelectOptions, errs map[string]string) []models.Field {

	id := models.NewInputFieldDisabled("id", "ID", incident.ID.String(), "text", "id", "", "", "")
	company := models.NewDropdown("company_id", "Company", companies, string(companies[0].Name), "", "/configuration-items-select", "#configuration_item_id")
	ci := models.NewDropdown("configuration_item_id", "Configuration Item", cis, string(cis[0].Name), "", "", "")
	state := models.NewDropdown("state", "State", states, string(states[0].Name), "", "", "")
	assignedTo := models.NewDropdown("assigned_to_id", "Assigned To", users, incident.AssignedTo.String(), "", "", "")
	shortDesc := models.NewInputField("short_description", "Short Description", incident.ShortDescription, "text", "short-description", "", "", "")
	desc := models.NewInputField("description", "Description", incident.Description, "text", "description", "", "", "")

	fields := []models.Field{
		&id,
		&company,
		&ci,
		&state,
		&assignedTo,
		&shortDesc,
		&desc,
	}

	// Apply the error text to the fields
	for i, e := range errs {
		for _, f := range fields {
			if i == f.GetID() {
				f.SetError(fmt.Sprintf("%s %s", f.GetLabel(), e))
			}

		}
	}

	return fields
}

func (cfg *ApiConfig) BuildIncidentsPage(r *http.Request, i models.Incident, u models.User, errs map[string]string) (templ.Component, error) {
	var companies models.SelectOptions
	err := cfg.GetCompaniesSelection(r, &companies)
	if err != nil {
		return nil, err
	}

	var cis models.SelectOptions
	err = cfg.GetCISelection(r, &cis)
	if err != nil {
		return nil, err
	}

	var users models.SelectOptions
	err = cfg.GetUsersSelection(r, &users)
	if err != nil {
		return nil, err
	}

	var states models.SelectOptions
	err = cfg.GetStatesSelection(r, &states)
	if err != nil {
		return nil, err
	}

	fields := MakeIncidentFields(i, companies, cis, states, users, errs)

	formData := models.NewFormData()
	// path := fmt.Sprintf("/incidents/%s", i.ID)
	path := "/incidents/add"
	cancelPath := "/incidents"
	form := models.NewIncidentForm("POST", path, cancelPath, companies, cis, states, users, i, formData)

	index := views.NewIncidentForm(form, fields)
	page := NewPage("Incidents - ADD", cfg, u)
	layout := views.BuildLayout(page, index)
	return layout, nil
}
