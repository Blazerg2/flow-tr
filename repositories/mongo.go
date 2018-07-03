package repositories

import (
	"os"
	"gopkg.in/mgo.v2"
)

func MongoRepo() *mgo.Database {
	var mongoURI = os.Getenv("MONGO_URI")
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		panic(err)
	}
	database := session.DB("test")
	return database
}
