package main

import (
	"log"

	"github.com/berrylradianh/ecowave-go/cmd/app"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	log.Println("Starting application...")
	route := app.StartApp()

	route.Start(":8080")
}
