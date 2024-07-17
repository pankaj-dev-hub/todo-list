package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConfig struct {
	Url        string
	Database   string
	Collection string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Url:        "mongodb://mongo-service:27017",
		Database:   "tododb",
		Collection: "todo",
	}

}

func ConnectDB(config *DBConfig) (*mongo.Database, error) {

	clientOption := options.Client().ApplyURI(config.Url)

	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(config.Database)

	return db, nil
}

func NewConnection() (*mongo.Database, *mongo.Collection) {

	configdb := NewDBConfig()

	db, err := ConnectDB(configdb)
	if err != nil {
		log.Fatal("Failed to connect db.", err)
	}

	collection := db.Collection(configdb.Collection)

	fmt.Println("Connected successfully.", collection)
	return db, collection
}
