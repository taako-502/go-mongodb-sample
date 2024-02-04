package infrastructures

import "go.mongodb.org/mongo-driver/mongo"

type MongoTransactionManager struct {
	Client *mongo.Client
}

func NewMongoTransactionManager(client *mongo.Client) *MongoTransactionManager {
	return &MongoTransactionManager{Client: client}
}

func (tm *MongoTransactionManager) StartSession() (mongo.Session, error) {
	return tm.Client.StartSession()
}
