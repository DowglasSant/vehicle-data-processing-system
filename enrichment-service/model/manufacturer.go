package model

type Manufacturer struct {
	ID                      int    `json:"id"`
	Name                    string `json:"name"`
	Headquarters            string `json:"headquarters"`
	DomesticOrInternational string `json:"domestic_or_international"`
	FoundationYear          int    `json:"foundation_year"`
}
