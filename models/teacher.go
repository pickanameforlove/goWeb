package models

const (
	CollectionTeacher = "teacher"
)

//教师编号tid、姓名name、性别sex、年龄age、院系名称dname
type Teacher struct {
	//ID    bson.ObjectId `json:"_id" bson:"_id"`
	Tid   string `json:"tid" bson:"tid"`
	Name  string `json:"name" bson:"name"`
	Sex   string `json:"sex" bson:"sex"`
	Age   int    `json:"age" bson:"age"`
	Dname string `json:"dname" bson:"dname"`
}
