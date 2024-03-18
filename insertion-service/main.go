package main

import (
	"context"
	"insertion-service/config"
	"insertion-service/consumer"
	"insertion-service/database"
	"insertion-service/repository"
	"insertion-service/service"
	"log"
)

func main() {
	log.Println("Starting the vehicle insertion service...")

	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbClient := database.ConnectMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer dbClient.Disconnect(context.Background())

	vehicleRepo := repository.NewVehicleRepository(dbClient, cfg.MongoDBName)
	vehicleSvc := service.NewVehicleService(vehicleRepo)

	consumer := consumer.NewConsumer(cfg.KafkaURL, cfg.KafkaTopic, cfg.GroupID, vehicleSvc)
	consumer.Consume()
}
