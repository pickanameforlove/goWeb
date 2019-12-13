package models

import (
	"time"
)

//type Type int

const (
	CollectionStudent = "student"
	//Int               Type = iota
	//String
)

//type Age struct {
//	Type   Type
//	IntVal int
//	StrVal string
//}
//
////实现json.Unmarshaller接口
//func (age *Age) UnmarshalJSON(value []byte) error {
//	if value[0] == '"' {
//		age.Type = String
//		return json.Unmarshal(value, &age.StrVal)
//	}
//	age.Type = Int
//	return json.Unmarshal(value, &age.IntVal)
//}
//
////实现josn.Marshaller接口
//func (age *Age) MarshalJSON() ([]byte, error) {
//	switch age.Type {
//	case Int:
//		return json.Marshal(age.IntVal)
//	case String:
//		return json.Marshal(age.StrVal)
//	default:
//		return []byte{}, fmt.Errorf("impossible Age.Type")
//	}
//}

//学生编号sid、姓名name、性别sex、年龄age、出生日期birthday、院系名称dname、班级class
type Student struct {
	//ID       bson.ObjectId `json:"_id" bson:"_id"`
	Sid      string    `json:"sid" bson:"sid"`
	Name     string    `json:"name" bson:"name"`
	Sex      string    `json:"sex" bson:"sex"`
	Age      int       `json:"age" bson:"age"`
	Birthday time.Time `json:"birthday" bson:"birthday"`
	Dname    string    `json:"dname" bson:"dname"`
	Class    string    `json:"class" bson:"class"`
}
