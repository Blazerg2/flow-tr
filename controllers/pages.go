package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"app/repositories"
)

func GetPagesList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request /pages")
	var results, err = repositories.AllPages()
	if err != nil {
		panic(err)
	}
	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
