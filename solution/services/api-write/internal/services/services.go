package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Service struct {
	// Add MongoDB and RabbitMQ connections here
	mongoClient *mongo.Client
	ramqConn    *amqp.Connection
}

func NewService(mongoClient *mongo.Client, ramqConn *amqp.Connection) *Service {
	return &Service{
		mongoClient: mongoClient,
		ramqConn:    ramqConn,
	}
}
