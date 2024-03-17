package repository

import (
	"database/sql"
	"log"

	"enrichment-service/model"

	_ "github.com/lib/pq"
)

type Repository interface {
	GetManufacturerByName(name string) (*model.Manufacturer, error)
	GetCityAndStateByName(cityName string) (*model.CityState, error)
}

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(connectionString string) *SQLRepository {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	return &SQLRepository{db: db}
}

func (r *SQLRepository) GetManufacturerByName(name string) (*model.Manufacturer, error) {
	var m model.Manufacturer
	query := `SELECT id, name, headquarters, domestic_or_international, foundation_year FROM vehicles.manufacturer WHERE name = $1`
	err := r.db.QueryRow(query, name).Scan(&m.ID, &m.Name, &m.Headquarters, &m.DomesticOrInternational, &m.FoundationYear)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *SQLRepository) GetCityAndStateByName(cityName string) (*model.CityState, error) {
	var cs model.CityState
	query := `SELECT c.id, c.name, s.id, s.name FROM vehicles.city c INNER JOIN vehicles.state s ON c.state_id = s.id WHERE c.name = $1`
	err := r.db.QueryRow(query, cityName).Scan(&cs.City.ID, &cs.City.Name, &cs.State.ID, &cs.State.Name)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}
