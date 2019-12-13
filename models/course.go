package models

const (
	CollectionCourse = "course"
)

//课程编号cid、课程名称name、先行课编号fcid、学分credit

type Course struct {
	//ID     bson.ObjectId `json:"_id" bson:"_id"`
	Cid    string `json:"cid" bson:"cid"`
	Name   string `json:"name" bson:"name"`
	Fcid   string `json:"fcid" bson:"fcid"`
	Credit string `json:"credit" bson:"credit"`
}
