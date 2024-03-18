package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	KafkaURL    string
	KafkaTopic  string
	MongoURI    string
	MongoDBName string
	GroupID     string
}

func LoadEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	return &Config{
		KafkaURL:    os.Getenv("KAFKA_URL"),
		KafkaTopic:  os.Getenv("KAFKA_TOPIC"),
		MongoURI:    os.Getenv("MONGO_URI"),
		MongoDBName: os.Getenv("MONGO_DB_NAME"),
		GroupID:     os.Getenv("GROUP_ID"),
	}, nil
}
