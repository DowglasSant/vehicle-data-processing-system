package kafkaconfig

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"enrichment-service/model"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer(writer *kafka.Writer) *KafkaProducer {
	return &KafkaProducer{Writer: writer}
}

func CreateWriter(kafkaURL, enrichedTopic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    enrichedTopic,
		Balancer: &kafka.LeastBytes{},
	})
}

func (kp *KafkaProducer) ProduceMessage(vehicle *model.Vehicle, topicType string) error {
	ctx := context.Background()
	enrichedData, err := json.Marshal(vehicle)
	if err != nil {
		log.Printf("[%s] Error marshalling enriched vehicle: %v", topicType, err)
		return err
	}

	err = kp.Writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(vehicle.LicensePlate),
		Value: enrichedData,
		Time:  time.Now(),
	})
	if err != nil {
		log.Printf("[%s] Failed to write message: %v", topicType, err)
		return err
	}
	log.Printf("[%s] Enriched message written to topic", topicType)
	return nil
}
