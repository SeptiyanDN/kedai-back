package businesses

import "time"

type BusinessesFormatter struct {
	ID            int       `json:"id"`
	Business_name string    `json:"business_name"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}

func FormatBusiness(Business Business) BusinessesFormatter {
	BusinessesFormatter := BusinessesFormatter{}
	BusinessesFormatter.ID = Business.ID
	BusinessesFormatter.Business_name = Business.Business_name
	return BusinessesFormatter
}

func FormatBusinesses(Businesses []Business) []BusinessesFormatter {
	BusinessesFormatter := []BusinessesFormatter{}
	for _, Business := range Businesses {
		BusinessFormatter := FormatBusiness(Business)
		BusinessesFormatter = append(BusinessesFormatter, BusinessFormatter)
	}
	return BusinessesFormatter
}
