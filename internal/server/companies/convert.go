package companies

import (
	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
)

func convert(company database.Company) data.Company {
	return data.Company{
		ID:        company.ID,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
		Name:      company.Name,
	}
}

func convertMany(companies []database.Company) []data.Company {
	var items []data.Company
	for _, item := range companies {
		items = append(items, convert(item))
	}
	return items
}

// func convertLatest(company database.GetCompaniesLatestRow) data.Company {
// 	return data.Company{
// 		ID:        company.ID,
// 		CreatedAt: company.CreatedAt,
// 		UpdatedAt: company.UpdatedAt,
// 		Name:      company.Name,
// 	}
// }

// func convertLatestMany(companies []database.GetCompaniesLatestRow) []data.Company {
// 	var items []data.Company
// 	for _, item := range companies {
// 		items = append(items, convertLatest(item))
// 	}
// 	return items
// }
