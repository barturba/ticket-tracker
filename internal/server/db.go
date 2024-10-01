package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/models"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) GetCompaniesSelection(r *http.Request, dst any) error {
	databaseCompanies, err := cfg.DB.GetCompanies(r.Context())
	if err != nil {
		return errors.New("couldn't find companies")
	}
	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)
	selectOptionsCompany := models.SelectOptions{}

	for _, company := range companies {
		selectOptionsCompany = append(selectOptionsCompany, models.NewSelectOption(company.Name, company.ID.String()))
	}
	*dst.(*models.SelectOptions) = selectOptionsCompany
	return nil
}

func (cfg *ApiConfig) GetStateSelection(dst any) {

	stateOptions := models.SelectOptions{}
	for _, so := range models.StateOptionsEnum {
		stateOptions = append(stateOptions, models.NewSelectOption(string(so), string(so)))
	}
	*dst.(*models.SelectOptions) = stateOptions
}

func NewIncident() models.Incident {
	return models.Incident{
		ID:                    uuid.New(),
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
		ShortDescription:      "",
		Description:           "",
		State:                 "",
		AssignedTo:            [16]byte{},
		AssignedToName:        "",
		ConfigurationItemID:   [16]byte{},
		ConfigurationItemName: "",
		CompanyID:             [16]byte{},
	}
}
