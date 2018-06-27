package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"os"
)

func main() {
	fmt.Println("Server started")
	router := NewRouter()
	server := http.ListenAndServe(GetPort(), router)
	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Roibs")
}

func GetPageVar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "toma tu jodida variable ")
	params := mux.Vars(r)
	fmt.Fprintf(w, params["pageVarId"])
}

func GetPagesList(w http.ResponseWriter, r *http.Request) {
	pages := Pages{
		Page{"la prueba 1", 1, 4},
		Page{"la prueba 2", 1, 4},
		Page{"la prueba 3", 1, 2},
		Page{"la prueba 4", 2, 4},
	}
	json.NewEncoder(w).Encode(pages)
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
