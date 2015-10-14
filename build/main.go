package main

import (
	"strconv"
	"github.com/amaniak/tesla"
)

func main() {

	tesla.Notice("Tesla Remote Interface Terminal")

	// Grab credentials
	cred := tesla.GetCredentials("./cr.conf")

	// connect to service
	service := tesla.Connect()

	// authorize and get token
	token := service.Authorize(cred)

	//Print token
	tesla.JQPrinter(token.ToJSON())

	//Print vehicles
	tesla.JQPrinter(service.Vehicle.All())

	// Ask
	idStr := tesla.Ask("Choose vehicle id: ")

	// Cast to int
	i, err := strconv.Atoi(idStr)

	//Honk Honk
	tesla.JQPrinter(service.Vehicle.HonkHonk(i))

}
