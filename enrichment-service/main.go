package main

import (
	"log"

	"enrichment-service/config"
	"enrichment-service/kafkaconfig"
	"enrichment-service/repository"
	"enrichment-service/service"
)

func main() {
	cfg := config.Load()

	log.Println("Starting enrichment service...")

	carReader := kafkaconfig.CreateReader(cfg.KafkaURL, cfg.TopicCar, cfg.GroupID)
	motobikeReader := kafkaconfig.CreateReader(cfg.KafkaURL, cfg.TopicMotorcycle, cfg.GroupID)

	defer carReader.Close()
	defer motobikeReader.Close()

	repo := repository.NewSQLRepository(cfg.DBConnectionString)

	writer := kafkaconfig.CreateWriter(cfg.KafkaURL, cfg.EnrichedTopic)
	defer writer.Close()

	kafkaProducer := kafkaconfig.NewKafkaProducer(writer)

	enrichmentService := service.NewVehicleEnrichmentService(repo, kafkaProducer)

	log.Println("Consumers and Producer initialized. Listening on topics:", cfg.TopicCar, "and", cfg.TopicMotorcycle)

	go kafkaconfig.ConsumeMessages(carReader, enrichmentService, "Car", cfg.EnrichedTopic)
	go kafkaconfig.ConsumeMessages(motobikeReader, enrichmentService, "Motobike", cfg.EnrichedTopic)

	select {}
}
