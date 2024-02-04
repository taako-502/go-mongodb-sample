package infrastructures

import "go.mongodb.org/mongo-driver/mongo"

type MongoTransactionManager struct {
	Client *mongo.Client
}

func NewMongoTransactionManager(client *mongo.Client) *MongoTransactionManager {
	return &MongoTransactionManager{Client: client}
}
