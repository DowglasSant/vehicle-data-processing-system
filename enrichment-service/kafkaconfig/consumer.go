package kafkaconfig

import (
	"context"
	"encoding/json"
	"log"

	"enrichment-service/model"
	"enrichment-service/service"

	"github.com/segmentio/kafka-go"
)

func CreateReader(kafkaURL, topic, groupID string) *kafka.Reader {
	log.Printf("Creating reader for topic: %s with group ID: %s\n", topic, groupID)
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
}

func ConsumeMessages(reader *kafka.Reader, service *service.VehicleEnrichmentService, topicType string, enrichedTopic string) {
	ctx := context.Background()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("[%s] Error reading message: %v", topicType, err)
			continue
		}

		log.Printf("[%s] Message received: %s", topicType, string(msg.Value))

		var vehicle model.Vehicle
		if err := json.Unmarshal(msg.Value, &vehicle); err != nil {
			log.Printf("[%s] Error unmarshalling message: %v", topicType, err)
			continue
		}

		err = service.EnrichVehicle(vehicle, topicType)
		if err != nil {
			log.Printf("[%s] Error enriching vehicle: %v", topicType, err)
			continue
		}
	}
}
