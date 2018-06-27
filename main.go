package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
)

func main() {
	fmt.Println("Server started")
	router := NewRouter()
	server := http.ListenAndServe(":8080", router)
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
