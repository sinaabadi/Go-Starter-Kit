package middlewares

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"starter-kit/config"
)

func ConnectToMongo() *mongo.Client {
	appConfig := config.GetConfig()
	clientOptions := options.Client().ApplyURI(appConfig.Get(`mongoUri`).(string))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf(`Error on connecting to mongo => %v`, err)
	}
	if pingErr := client.Ping(context.TODO(), nil); pingErr != nil {
		log.Fatalf(`Error on connecting to mongo => %v`, pingErr)
	}
	log.Println(`Connected to mongo`)
	return client
}
