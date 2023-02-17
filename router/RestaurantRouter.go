package router

import (
	"github.com/faisalmoinuddin99/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/v1/api/restaurant", controller.CreateRestaurant).Methods("POST")
	router.HandleFunc("/v1/api/restaurants", controller.GetAllRestaurants).Methods("GET")
	router.HandleFunc("/v1/api/restaurant/{id}", controller.GetRestaurant).Methods("GET")
	router.HandleFunc("/v1/api/restaurant/{id}", controller.DeleteOneRestaurant).Methods("DELETE")

	return router
}
