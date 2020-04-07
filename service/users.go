package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"techTask/models"
)

//CONNSTRING Database location
const CONNSTRING = "mongodb://127.0.0.1:27017"

//DBNAME database name
const DBNAME = "coinCover"

//COLLNAME Collection name
const COLLNAME = "users"

var db *mongo.Database

//Connect handles the connection to the database
func init() {
	clientOptions := options.Client().ApplyURI(CONNSTRING)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Problem connecting to DB. Shutting down")
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
		fmt.Println("Could not insert user")
		fmt.Println(err)
	}

}

//GetAllUsers retrieves all entries in the users collection of the database
func GetAllUsers() []models.User {

	cur, err := db.Collection(COLLNAME).Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var elements []models.User
	var elem models.User

	//For loop gets the next result from the cursor
	for cur.Next(context.TODO()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return elements

}
