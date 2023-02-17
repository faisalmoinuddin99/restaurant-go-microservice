package controller

import (
	"encoding/json"
	"net/http"

	"github.com/faisalmoinuddin99/model"
	"github.com/faisalmoinuddin99/service"
	"github.com/gorilla/mux"
)

// POST
func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var restaurant model.Restaurant
	err := json.NewDecoder(r.Body).Decode(&restaurant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	service.InsertOneRestaurant(restaurant)
	json.NewEncoder(w).Encode(restaurant)
}

// GET ALL 
func GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	allRestaurant := service.GetAllRestaurant()

	json.NewEncoder(w).Encode(allRestaurant)
}

// GET SINGLE
func GetRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	restaurant := service.GetRestaurantById(params["id"])

	json.NewEncoder(w).Encode(restaurant)
}

// DELETE SINGLE
func DeleteOneRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	message := service.DeleteRestaurantById(params["id"])

	json.NewEncoder(w).Encode(message)
}
