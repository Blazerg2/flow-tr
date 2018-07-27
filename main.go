package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"os"
	"gopkg.in/mgo.v2"
	"crypto/tls"
	"net"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"strconv"
)

func main() {
	fmt.Println("Server started")
	router := NewRouter()
	server := http.ListenAndServe(GetPort(), router)
	log.Fatal(server)
}

var mongoURI = os.Getenv("ATLAS_TEST_G")
var collection = getSession().DB("test").C("pages")

func getSession() *mgo.Session {
	dialInfo, err := mgo.ParseURL(mongoURI)
	if (err != nil) {
		panic(err)
	}
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if (err != nil) {
		panic(err)
	}
	log.Println("MongoDB cluster connected")
	return session
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my underworld")
}

func GetPage(w http.ResponseWriter, r *http.Request) {
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
