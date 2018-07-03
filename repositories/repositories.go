package repositories

var collection = MongoRepo().C("pages")

// AllPages from mongo repo
func AllPages() ([]Page, error) {
	var results []Page
	err := collection.Find(nil).All(&results)
	// if err != nil {
	// 	panic(err)
	// }
	return results, err
}

// SavePage save in pages conllection
func SavePage(pageData Page) (Page, error) {
	err := collection.Insert(pageData)
	return pageData, err
}

// Page struct
type Page struct {
	Text      string `json:"text"`
	InStates  int    `json:"intStates"`
	OutStates int    `json:"outStates"`
}

// Pages struct
type Pages []Page

// var mongoURI = os.Getenv("MONGO_URI")
// var collection = getSession().DB(os.Getenv("MONGO_DATABASE")).C("pages")

// func getSession() *mgo.Session {
// 	dialInfo, err := mgo.ParseURL(mongoURI)
// 	if (err != nil) {
// 		panic(err)
// 	}
// 	tlsConfig := &tls.Config{}
// 	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
// 		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
// 		return conn, err
// 	}
// 	session, err := mgo.DialWithInfo(dialInfo)
// 	if (err != nil) {
// 		panic(err)
// 	}
// 	log.Println("MongoDB cluster connected")
// 	return session
// }
