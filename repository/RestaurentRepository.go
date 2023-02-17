package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://faisal25:faisalfacebook@restaurant.oseq8oz.mongodb.net/?retryWrites=true&w=majority"
const dbName = "restaurantDB"
const collectionName = "restaurantList"

// MOST IMPORTANT
var collection *mongo.Collection

// connnect with mongodb
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection Success")

	collection = client.Database(dbName).Collection(collectionName)

	// collection instance
	fmt.Println("Collections instance is ready")
}
