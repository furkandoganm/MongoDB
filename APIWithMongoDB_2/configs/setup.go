package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	//URI tarafında belirtilen bir dağıtıma bağlanmak için istemci oluşturduk

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Connect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("apiGolang").Collection(collectionName)
}
