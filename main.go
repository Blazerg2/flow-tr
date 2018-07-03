package main

import (
	"flow-tr/utils"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := GetPort()
	router := utils.SimpleRouter(AppRoutes())
	server := http.ListenAndServe(port, router)
	fmt.Println("Server started")
	log.Fatal(server)
}

// GetPort the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3001"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
