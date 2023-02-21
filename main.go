package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faisalmoinuddin99/configuration"
	"github.com/faisalmoinuddin99/router"
)

func main() {
	fmt.Println("Welcome to Restuarent Backend Project")

	r := router.Router()

	// configuration.Config()
	configuration.EurekClientConfig()
	/*
		consul agent -server -bootstrap-expect=1 -data-dir=consul-data -ui -bind=192.168.1.103
										OR
		consul agent -dev
	*/
	fmt.Println("Server is getting started ....")
	log.Fatal(http.ListenAndServe(":8085", r))
	fmt.Println("Listening at port 8085 ...")
}
