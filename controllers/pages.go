package controllers

import (
	"encoding/json"
	"flow-tr/repositories"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index action
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my underworld")
}

// GetPagesList action
func GetPagesList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request /pages")
	var results, err = repositories.AllPages()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

// GetPageVar action
func GetPageVar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "toma tu jodida variable ")
	params := mux.Vars(r)
	fmt.Fprintf(w, params["pageVarId"])
}

// PostPagesList action
func PostPagesList(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var pageData repositories.Page
	err := decoder.Decode(&pageData)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	pageData, err = repositories.SavePage(pageData)

	if err != nil {
		w.WriteHeader(500)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(pageData)
}
