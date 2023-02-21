package configuration

import (
	"log"

	"github.com/hashicorp/consul/api"
)

func Config() {
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      "restaurant-service-8085",
		Name:    "restaurant-service",
		Port:    8085,
		Address: "localhost",
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}
}
