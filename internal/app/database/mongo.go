package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func NewDatabaseConn() *mongo.Database {
	host := os.Getenv("DB_HOST")
	/*
		dbUsername := os.Getenv("DB_USERNAME")
		dbPassword := os.Getenv("DB_PASSWORD")
	*/
	database := os.Getenv("DB_NAME")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(host))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())
	fmt.Println("DB:", database)
	db := client.Database(database)

	return db
}
