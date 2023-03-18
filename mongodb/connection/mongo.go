package connection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Context context.Context
	MongoDB *mongo.Client
}

type IMongoDB interface {
	Disconnect()
}

func NewMongoDB(connectionURL string) *MongoDB {
	option := options.Client().ApplyURI(connectionURL)
	client, err0 := mongo.Connect(context.Background(), option)

	if err0 != nil {
		log.Fatal(err0)
	}

	return &MongoDB{
		Context: context.Background(),
		MongoDB: client,
	}
}

func (ds MongoDB) Disconnect() {
	ds.MongoDB.Disconnect(ds.Context)
}
