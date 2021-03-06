package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net"
	"net/http"
	"os"
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
	if err != nil {
		panic(err)
	}
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	log.Println("MongoDB cluster connected")
	return session
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my underworld") //unhandled error
}

func GetPageVar(w http.ResponseWriter, r *http.Request) {
	//TODO stop insulting the user and implement this LOL
	fmt.Fprintf(w, "Take your stupid variable") //unhandled error
	params := mux.Vars(r)
	fmt.Fprintf(w, params["pageVarId"]) //unhandled error
}

func GetPagesList(w http.ResponseWriter, r *http.Request) {
	var results []Page
	err := collection.Find(nil).All(&results)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results) //unhandled error
}

func PostPagesList(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var page_data Page
	err := decoder.Decode(&page_data)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close() //unhandled error

	err = collection.Insert(page_data)

	if err != nil {
		w.WriteHeader(500)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(page_data) //unhandled error
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
