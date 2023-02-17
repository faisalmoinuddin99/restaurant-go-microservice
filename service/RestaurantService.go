package service

import (
	"context"
	"fmt"
	"log"

	"github.com/faisalmoinuddin99/model"
	"github.com/faisalmoinuddin99/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MOST IMPORTANT
var collection = repository.Collection

func insertOneRestaurant(restaurant model.Restuarent) {

	inserted, err := collection.InsertOne(context.Background(), restaurant)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data Inserted in restaurant db with id: ", inserted.InsertedID.(primitive.ObjectID))
}


func getAllRestaurant() []model.Restuarent {
	cursor, err := collection.Find(context.Background(),bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var restaurants []model.Restuarent
	for cursor.Next(context.Background()){
		var restaurant model.Restuarent
			err := cursor.Decode(&restaurant)
			if err != nil {
				log.Fatal(err)
			}
			restaurants = append(restaurants, restaurant)
	}
	defer cursor.Close(context.Background())
	return restaurants
}