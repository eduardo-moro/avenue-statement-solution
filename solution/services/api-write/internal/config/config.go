package config

import (
	"log"
	"os"
)

type Config struct {
	MongoURI    string
	RabbitMQURI string
	Port        string
}

func LoadConfig() *Config {
	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		log.Fatal("MONGO_URL environment variable not set")
	}

	rabbitMQURI := os.Getenv("RABBITMQ_URL")
	if rabbitMQURI == "" {
		log.Fatal("RABBITMQ_URL environment variable not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		MongoURI:    mongoURI,
		RabbitMQURI: rabbitMQURI,
		Port:        port,
	}
}
