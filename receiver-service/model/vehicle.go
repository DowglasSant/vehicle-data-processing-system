package model

type Vehicle struct {
	Type             int    `json:"type"`
	LicensePlate     string `json:"license_plate"`
	Year             int    `json:"year"`
	OwnerCPF         string `json:"owner_cpf"`
	RegistrationCity string `json:"registration_city"`
	Color            string `json:"color"`
	Brand            string `json:"brand"`
	Model            string `json:"model"`
}
