package main

func main() {
	route := startApp()

	route.Start(":8080")
}