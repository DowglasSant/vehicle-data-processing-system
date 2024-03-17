package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	KafkaURL           string `envconfig:"KAFKA_URL"`
	GroupID            string `envconfig:"KAFKA_GROUP_ID"`
	TopicCar           string `envconfig:"CAR_TOPIC"`
	TopicMotorcycle    string `envconfig:"MOTOBYKE_TOPIC"`
	EnrichedTopic      string `envconfig:"ENRICHED_TOPIC"`
	DBConnectionString string `envconfig:"DB_CONNECTION_STRING"`
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal("Failed to load environment variables: ", err)
	}

	checkRequiredEnvVars(&cfg)

	return cfg
}

func checkRequiredEnvVars(cfg *Config) {
	var missingVars []string
	if cfg.KafkaURL == "" {
		missingVars = append(missingVars, "KAFKA_URL")
	}
	if cfg.GroupID == "" {
		missingVars = append(missingVars, "KAFKA_GROUP_ID")
	}
	if cfg.DBConnectionString == "" {
		missingVars = append(missingVars, "CONNECTION_STRING")
	}

	if len(missingVars) > 0 {
		log.Fatalf("Missing required environment variables: %v", missingVars)
	}
}
