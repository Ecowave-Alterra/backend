package main

import (
	"log"

	"github.com/berrylradianh/ecowave-go/cmd/app"
)

func main() {
	log.Println("Starting application...")
	route := app.StartApp()

	route.Start(":8080")
}
