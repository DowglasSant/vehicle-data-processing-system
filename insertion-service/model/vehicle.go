package model

import "time"

type Vehicle struct {
	Type             int          `json:"type"`
	LicensePlate     string       `json:"license_plate"`
	Year             int          `json:"year"`
	OwnerCPF         string       `json:"owner_cpf"`
	RegistrationCity string       `json:"registration_city"`
	Color            string       `json:"color"`
	Brand            string       `json:"brand"`
	Model            string       `json:"model"`
	Manufacturer     Manufacturer `json:"manufacturer_info"`
	State            State        `json:"state_info"`
	InsertionTime    time.Time    `json:"insertion_time"`
}
