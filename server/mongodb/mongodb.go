package mongodb

import (
	"context"
	"latifa/config"
	log2 "log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient() (*mongo.Client, error) {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(config.Get().MongoDBUri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log2.Fatal("mongodb/client: failed to ping mongodb cluster: ", err)
	}

	/*defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log2.Fatal("mongodb/client: failed to disconnect from mongodb: ", err)
		}
	}()*/

	return client, nil
}
