package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"avenue-golang/services/api-write/internal/config"
	"avenue-golang/services/api-write/internal/handlers"
	"avenue-golang/services/api-write/internal/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	// Connect to RabbitMQ
	ramqConn, err := amqp.Dial(cfg.RabbitMQURI)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer ramqConn.Close()

	log.Println("Connected to MongoDB and RabbitMQ")

	svc := services.NewService(client, ramqConn)
	h := handlers.NewHandler(svc)

	http.HandleFunc("/health", h.HealthCheck)
	http.HandleFunc("/events", h.EventsHandler)

	log.Printf("API Write Service listening on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
