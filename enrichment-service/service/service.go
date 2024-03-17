package service

import (
	"enrichment-service/model"
	"enrichment-service/repository"
)

type VehicleEnrichmentService struct {
	repo        repository.Repository
	messageProd MessageProducer
}

func NewVehicleEnrichmentService(repo repository.Repository, mp MessageProducer) *VehicleEnrichmentService {
	return &VehicleEnrichmentService{repo: repo, messageProd: mp}
}

func (s *VehicleEnrichmentService) EnrichVehicle(vehicle model.Vehicle, topicType string) error {
	manufacturer, err := s.repo.GetManufacturerByName(vehicle.Brand)
	if err != nil {
		return err
	}
	vehicle.ManufacturerInfo = *manufacturer

	cityState, err := s.repo.GetCityAndStateByName(vehicle.RegistrationCity)
	if err != nil {
		return err
	}
	vehicle.StateInfo = cityState.State

	return s.messageProd.ProduceMessage(&vehicle, topicType)
}
