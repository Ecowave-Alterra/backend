package main

import (
	"github.com/berrylradianh/ecowave-go/cmd/app"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../.env")
}

func main() {
	route := app.StartApp()

	route.Start(":8080")
}
