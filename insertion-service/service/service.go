package service

import (
	"insertion-service/model"
	"insertion-service/repository"
	"time"
)

type VehicleService struct {
	repo *repository.VehicleRepository
}

func NewVehicleService(repo *repository.VehicleRepository) *VehicleService {
	return &VehicleService{repo: repo}
}

func (svc *VehicleService) InsertVehicle(vehicle model.Vehicle) error {
	vehicle.InsertionTime = time.Now()
	return svc.repo.InsertVehicle(vehicle)
}
