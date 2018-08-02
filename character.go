package main

import (
	"net/http"
	"encoding/json"
)

type Character struct {
	Character []string `json:"character"`
}

func GetCharacters(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	char := Character{}
	var tempChar []string

	err := collection.Find(nil).Distinct("character", &tempChar)
	char.Character = tempChar
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(char)
}
