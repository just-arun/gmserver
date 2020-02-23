package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const (
	Database = "GoTest"
)

type collection struct {
	Users string
	Posts string
}

func CollectionList() collection {
	return collection{
		Users: "user",
		Posts: "posts",
	}
}

func Init() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func CollectionFun(cli *mongo.Client, collectionName string) (*mongo.Collection, context.Context) {
	collection := cli.Database(Database).Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return collection, ctx
}
