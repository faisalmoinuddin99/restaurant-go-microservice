package configuration

import (
	"net"

	"github.com/ArthurHlt/go-eureka-client/eureka"
)

func EurekClientConfig() {
	// Get the IP address of the local machine
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	var ipAddress string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ipAddress = ipnet.IP.String()
			break
		}
	}

	// Initialize the Eureka client with server URL
	client := eureka.NewClient([]string{"http://localhost:8761/eureka"})
	// Create the instance information for the GoLang service

	// eureka.NewInstanceInfo(appName, ipAddr, hostname, port, securePort, ttl, secure)
	instance := eureka.NewInstanceInfo(
		"localhost",
		"RESTAURANT-SERVICE",
		ipAddress,
		8085,
		30,
		false,
	)

	// Register the instance with the Eureka server
	client.RegisterInstance("RESTAURANT-SERVICE", instance)
	client.SendHeartbeat(instance.App, instance.HostName)

}
