package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faisalmoinuddin99/router"
)

func main() {
	fmt.Println("Welcome to Restuarent Backend Project")

	r := router.Router()
	fmt.Println("Server is getting started ....")
	log.Fatal(http.ListenAndServe(":8080", r))
	log.Println("Listening at port 8080 ...")
}
