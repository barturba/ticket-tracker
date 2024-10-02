package server

import "github.com/barturba/ticket-tracker/models"

func MakeIncidentFields(incident models.Incident, companies, cis, states models.SelectOptions) []models.Field {
	id := models.NewInputFieldDisabled("id", "id", incident.ID.String(), "text", "id", "", "", "")
	company := models.NewDropdown("company_id", "Company", companies, string(companies[0].Name), "", "/configuration-items-select", "#configuration_item_id")
	ci := models.NewDropdown("configuration_item_id", "Configuration Item", cis, string(cis[0].Name), "", "", "")
	state := models.NewDropdown("state", "State", states, string(states[0].Name), "", "", "")
	shortDesc := models.NewInputField("short_description", "Short Description", incident.ShortDescription, "text", "short-description", "", "", "")
	desc := models.NewInputField("description", "Description", incident.Description, "text", "description", "", "", "")

	fields := []models.Field{
		&id,
		&company,
		&ci,
		&state,
		&shortDesc,
		&desc,
	}
	return fields
}
