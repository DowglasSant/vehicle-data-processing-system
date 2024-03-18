package repository

import (
	"context"
	"insertion-service/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type VehicleRepository struct {
	client *mongo.Client
	dbName string
}

func NewVehicleRepository(client *mongo.Client, dbName string) *VehicleRepository {
	return &VehicleRepository{
		client: client,
		dbName: dbName,
	}
}

func (repo *VehicleRepository) InsertVehicle(vehicle model.Vehicle) error {
	collectionName := "cars"
	if vehicle.Type == 2 {
		collectionName = "motos"
	}

	collection := repo.client.Database(repo.dbName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), vehicle)
	return err
}
