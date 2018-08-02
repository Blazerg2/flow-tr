package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"strconv"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

type Page struct {
	Text      string     `json:"text"`
	Instate   int        `json:"instate"`
	Character string     `json:"character"`
	Decisions []Decision `json:"decisions"`
	IsFinal   bool       `json:"isFinal"`
}

type Decision struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	page := Page{}
	//keys, ok := r.URL.Query()["character"]
	keys := r.URL.Query()
	key := keys.Get("character")

	if key == "" {
		log.Println("Url Param 'character' is missing")
		return
	}

	vars := mux.Vars(r)
	instate, err := strconv.Atoi(vars["pageId"])
	if err != nil {
		return
	}

	err = collection.Find(bson.M{"instate": instate, "character": key}).One(&page)

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(page)
}
func GetPagesList(w http.ResponseWriter, r *http.Request) {
	var results []Page
	err := collection.Find(nil).All(&results)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
func PostPagesList(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var page_data Page
	err := decoder.Decode(&page_data)
	if (err != nil) {
		panic(err)
	}
	defer r.Body.Close()

	err = collection.Insert(page_data)

	if (err != nil) {
		w.WriteHeader(500)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(page_data)
}