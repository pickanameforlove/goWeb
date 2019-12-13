package models

const (
	CollectionStudent_Course = "student_course"
)

//学号sid、课程号cid、成绩score、教师编号tid
type Student_Course struct {
	//ID  bson.ObjectId `json:"_id" bson:"_id"`
	Sid   string `json:"sid" bson:"sid"`
	Cid   string `json:"cid" bson;"cid"`
	Tid   string `json:"tid" bson:"tid"`
	Score int    `json:"score" bson:"score"`
}
