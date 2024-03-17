package main

import (
	"context"
	"encoding/json"
	"insertion-service/database"
	"insertion-service/model"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func main() {
	log.Println("Starting the vehicle insertion service...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	kafkaURL := os.Getenv("KAFKA_URL")
	topic := os.Getenv("KAFKA_TOPIC")
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	groupId := os.Getenv("GROUP_ID")

	log.Printf("Connecting to MongoDB at %s...", mongoURI)
	client := database.ConnectMongoDB(mongoURI)
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
		log.Println("Disconnected from MongoDB.")
	}()
	log.Println("Successfully connected to MongoDB.")

	db := client.Database(dbName)

	log.Printf("Subscribing to Kafka topic %s at %s...", topic, kafkaURL)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaURL},
		GroupID: groupId,
		Topic:   topic,
	})
	defer func() {
		if err := r.Close(); err != nil {
			log.Fatalf("Failed to close Kafka reader: %v", err)
		}
		log.Println("Kafka reader closed.")
	}()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		log.Printf("Received message: %s", string(m.Value))

		var vehicle model.Vehicle
		if err := json.Unmarshal(m.Value, &vehicle); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		collectionName := "cars"
		if vehicle.Type == 2 {
			collectionName = "motos"
		}

		vehicle.InsertionTime = time.Now()

		log.Printf("Inserting vehicle into %s collection...", collectionName)
		collection := db.Collection(collectionName)
		_, err = collection.InsertOne(context.Background(), vehicle)
		if err != nil {
			log.Printf("Error inserting document into %s: %v", collectionName, err)
			continue
		}

		log.Printf("Successfully inserted vehicle into %s collection.", collectionName)
	}
}
