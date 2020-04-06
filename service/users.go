package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"techTask/models"
)

//CONNSTRING Database location
const CONNSTRING = "mongodb://localhost:27017"

//DBNAME database name
const DBNAME = "coinCover"

//COLLNAME Collection name
const COLLNAME = "users"

var db *mongo.Database

//Connect handles the connection to the database
func Connect() {
	clientOptions := options.Client().ApplyURI(CONNSTRING)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(DBNAME)

	fmt.Println("Connected to DB")

}

//InsertUser inserts user into database
func InsertUser(user models.User) {
	fmt.Println(user)

	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

}
