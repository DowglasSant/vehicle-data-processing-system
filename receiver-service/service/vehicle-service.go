package service

import (
	"encoding/json"
	"log"
	"os"
	"receiver-service/kafka"
	"receiver-service/model"
)

func ProcessVehicleData(vehicle model.Vehicle) error {
	writer := kafka.LoadKafkaConfig()
	publisher := kafka.NewKafkaPublisher(writer)

	vehicleBytes, err := json.Marshal(vehicle)
	if err != nil {
		log.Printf("Error when serializing vehicle: %v\n", err)
		return err
	}

	var topic string
	if vehicle.Type == 1 {
		topic = os.Getenv("CAR_TOPIC")
	} else if vehicle.Type == 2 {
		topic = os.Getenv("MOTOBYKE_TOPIC")
	} else {
		log.Println("Vehicle with invalid type!")
		return err
	}

	err = publisher.Publish(topic, vehicle.LicensePlate, vehicleBytes)
	if err != nil {
		log.Printf("Error trying to publish message: %v\n", err)
		return err
	}

	return nil
}
