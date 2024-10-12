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

func convertRowsAndMetadata(rows []database.GetCompaniesRow, filters data.Filters) ([]data.Company, data.Metadata) {
	var output []data.Company
	var totalRecords int64 = 0
	for _, row := range rows {
		outputRow := convertRowAndCount(row, &totalRecords)
		output = append(output, outputRow)
	}
	metadata := data.CalculateMetadata(int(totalRecords), filters.Page, filters.PageSize)
	return output, metadata
}

func convertRowAndCount(row database.GetCompaniesRow, count *int64) data.Company {
	outputRow := data.Company{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Name:      row.Name,
	}
	count = &row.Count

	return outputRow
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
