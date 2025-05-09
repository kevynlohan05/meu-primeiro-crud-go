package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL           = "MONGODB_URL"
	MONGODB_USER_DATABASE = "MONGODB_USER_DATABASE"
)

func NewMongoDbConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_USER_DATABASE)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_database), nil
}
