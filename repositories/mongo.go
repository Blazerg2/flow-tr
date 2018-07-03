package repositories

import (
	"os"

	"gopkg.in/mgo.v2"
)

// MongoRepo get mongo connection
func MongoRepo() *mgo.Database {
	var mongoURI = os.Getenv("MONGO_URI")
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		panic(err)
	}
	database := session.DB("test")
	return database
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
