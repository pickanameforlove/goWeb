package db

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	Mongo   *mgo.DialInfo
)

const (
	MongoDBUrl = "mongodb://localhost:27017/user201700301079"
)

var (
	Db *mgo.Database
)

func init() {
	//os.Getenv得到环境变量.
	uri := os.Getenv("MONGODB_URL")
	if len(uri) == 0 {
		uri = MongoDBUrl
	}
	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Println("Can't connect to mongo,go error %v", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("connect to ", uri)
	Session = s
	Mongo = mongo
	DBint()

}
func DBint() {
	Db = Session.DB(Mongo.Database)
}
