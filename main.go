package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"gopkg.in/mgo.v2"
	"crypto/tls"
	"net"
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
