package consumer

import (
	"context"
	"encoding/json"
	"insertion-service/model"
	"insertion-service/service"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader  *kafka.Reader
	service *service.VehicleService
	topic   string
}

func NewConsumer(kafkaURL, topic, groupID string, svc *service.VehicleService) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaURL},
		GroupID: groupID,
		Topic:   topic,
	})
	return &Consumer{reader: r, service: svc, topic: topic}
}

func (c *Consumer) Consume() {
	log.Printf("Iniciando o consumo do tópico Kafka: %s", c.topic)
	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Erro ao ler mensagem do tópico %s: %v", c.topic, err)
			continue
		}

		log.Printf("Mensagem lida do tópico %s: %s", c.topic, string(m.Value)) // Apresente a mensagem lida

		var vehicle model.Vehicle
		if err := json.Unmarshal(m.Value, &vehicle); err != nil {
			log.Printf("Erro ao deserializar a mensagem do tópico %s: %v", c.topic, err)
			continue
		}

		if err := c.service.InsertVehicle(vehicle); err != nil {
			log.Printf("Erro ao inserir veículo do tópico %s: %v", c.topic, err)
			continue
		}

		log.Println("Veículo inserido com sucesso.")
	}
}
