package database

import (
	"context"
	"fiber-boilerplate/pkg/config"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	*mongo.Database
}

var db = &DB{}

func (db *DB) connect(config *config.DB) error {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	dbURI := fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority&appName=%s", config.Prefix, config.Username, config.Password, config.Host, config.Name)

	opts := options.Client().ApplyURI(dbURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		return err
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}, {"$db", "admin"}}).Err(); err != nil {
		return err
	}

	db.Database = client.Database(config.Name)

	fmt.Println("Successfully connected to MongoDB!")

	return nil
}

func GetDB() *DB {
	return db
}

func ConnectDB() error {
	return db.connect(config.DBCfg())
}
