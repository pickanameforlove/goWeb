package models

const (
	CollectionTeacher_Course = "teacher_course"
)

//教师编号tid、课程编号cid
type Teacher_Course struct {
	//ID  bson.ObjectId `json:"_id" bson:"_id"`
	Tid string `json:"tid" bson:"tid"`
	Cid string `json:"cid" bson:"cid"`
}
