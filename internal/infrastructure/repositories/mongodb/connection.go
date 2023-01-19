package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUser = "admin"
	dbPass = "secret"
	dbName = "shop_db"
)

func GetConnection() (*mongo.Client, error) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPass+"@localhost:27017"))
	if err != nil {
		panic("error connecting to DB: " + err.Error())
	}

	return client, nil
}
