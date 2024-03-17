package kafka

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	writer *kafka.Writer
}

func NewKafkaPublisher(writer *kafka.Writer) *KafkaPublisher {
	return &KafkaPublisher{writer: writer}
}

func NewKafkaWriter(kafkaURL string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Balancer: &kafka.LeastBytes{},
	}
}

func LoadKafkaConfig() *kafka.Writer {
	kafkaURL := os.Getenv("KAFKA_URL")
	if kafkaURL == "" {
		log.Fatal("KAFKA_URL environment variable is not set")
	}

	return NewKafkaWriter(kafkaURL)
}

func (kp *KafkaPublisher) Publish(topic string, key string, message []byte) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: message,
		Topic: topic,
	}
	return kp.writer.WriteMessages(context.Background(), msg)
}
