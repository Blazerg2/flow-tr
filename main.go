package main

import (
	"fmt"
	"net/http"
	// "log"
	"os"
	"app/controllers"
	// "crypto/tls"
	// "net"
)

func main() {
	// router := NewRouter()
	port := ":3001"
	fmt.Println("Server start %s", port )
	http.HandleFunc("/pages", controllers.GetPagesList)
	// server := http.ListenAndServe(port, nil)
	fmt.Println("Server start %s", port )
	if err := http.ListenAndServe(port, nil); err != nil {
    panic(err)
  }
	// log.Fatal(server)
}

var routes = Routes{
	// Route{
	// 	"Index",
	// 	"GET",
	// 	"/",
	// 	Index,
	// },
	// Route{
	// 	"GetPageVar",
	// 	"GET",
	// 	"/pagevar/{pageVarId}",
	// 	GetPageVar,
	// },
	Route{
		"GetPagesList",
		"GET",
		"/pages",
		controllers.GetPagesList,
	},
	// Route{
	// 	"PostPagesList",
	// 	"POST",
	// 	"/pages",
	// 	PostPagesList,
	// },
}



// func Index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello my underworld")
// }

// func GetPageVar(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "toma tu jodida variable ")
// 	params := mux.Vars(r)
// 	fmt.Fprintf(w, params["pageVarId"])
// }



// func PostPagesList(w http.ResponseWriter, r *http.Request) {
// 	decoder := json.NewDecoder(r.Body)

// 	var page_data Page
// 	err := decoder.Decode(&page_data)
// 	if (err != nil) {
// 		panic(err)
// 	}
// 	defer r.Body.Close()

// 	err = collection.Insert(page_data)

// 	if (err != nil) {
// 		w.WriteHeader(500)
// 		panic(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(200)
// 	json.NewEncoder(w).Encode(page_data)
// }

// // Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

// func MongoRepo() {
// 	// 	dialInfo, err := mgo.ParseURL(mongoURI)
// 	// 	if (err != nil) {
// 	// 		panic(err)
// 	// 	}
// 	// 	tlsConfig := &tls.Config{}
// 	// 	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
// 	// 		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
// 	// 		return conn, err
// 	// 	}
// 	// 	session, err := mgo.DialWithInfo(dialInfo)
// 	// 	if (err != nil) {
// 	// 		panic(err)
// 	// 	}
// 	// 	log.Println("MongoDB cluster connected")
// 	// 	return session

// 	var mongoURI = os.Getenv("MONGO_URI")
// 	session, err := mgo.Dial(mongoURI)
// 	if errr != nil {
// 		panic(errr)
// 	}
// 	database := session.DB("test")
// 	return database
// }
