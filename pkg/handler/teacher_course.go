package handler

import (
	mdb "MongodbPro/db"
	"MongodbPro/models"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func GetTid(cid string) (tid string, err error) {
	db := mdb.Db
	var tc models.Teacher_Course
	err = db.C(models.CollectionTeacher_Course).Find(bson.M{"cid": cid}).One(&tc)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	tid = tc.Tid
	return
}
