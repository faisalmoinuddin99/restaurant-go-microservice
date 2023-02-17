package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/faisalmoinuddin99/model"
	"github.com/faisalmoinuddin99/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MOST IMPORTANT
var collection = repository.Collection

func InsertOneRestaurant(restaurant model.Restaurant) {
	inserted, err := collection.InsertOne(context.Background(), restaurant)
	if err != nil {
		log.Fatal(err)
	}

	id, ok := inserted.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("Failed to get the inserted ID")
	}

	restaurant.ID = id
	jsonBytes, err := json.Marshal(restaurant)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data Inserted in restaurant db with id:", id)
	fmt.Println("Inserted data:", string(jsonBytes))
}

func GetAllRestaurant() []model.Restaurant {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var restaurants []model.Restaurant
	for cursor.Next(context.Background()) {
		var restaurant model.Restaurant
		err := cursor.Decode(&restaurant)
		if err != nil {
			log.Fatal(err)
		}
		restaurants = append(restaurants, restaurant)
	}
	defer cursor.Close(context.Background())
	return restaurants
}

func GetRestaurantById(restaurantId string) *model.Restaurant {

	// Create an empty Restaurant struct to hold the query result
	restaurant := &model.Restaurant{}

	objID, err := primitive.ObjectIDFromHex(restaurantId)
	if err != nil {
		log.Fatal(err)
	}

	// Build a filter to match the input ID
	filter := bson.M{"_id": objID}

	// Execute the query and store the result in the restaurant variable
	collection.FindOne(context.Background(), filter).Decode(&restaurant)

	// Return the restaurant object and no error
	return restaurant
}
func DeleteRestaurantById(restaurantId string) string {

	objID, err := primitive.ObjectIDFromHex(restaurantId)

	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id":objID}
	collection.DeleteOne(context.Background(),filter)

	return "ID: " + objID.String() + " Deleted Successfully...."
}