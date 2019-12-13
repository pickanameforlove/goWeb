package handler

import (
	mdb "MongodbPro/db"
	"MongodbPro/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2/bson"
)

func GetCourseBysid(sid string) (scs []models.Student_Course, err error) {
	db := mdb.Db
	err = db.C(models.CollectionStudent_Course).Find(bson.M{"sid": sid}).All(&scs)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	return
}

func InsertSC(c *gin.Context) {
	d, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	var sc models.Student_Course
	err1 := json.Unmarshal(d, &sc)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err1.Error(),
		})
		return
	}

	scs, err4 := GetCourseBysid(sc.Sid)
	if err4 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err4.Error(),
		})
		return
	}
	for i := 0; i < len(scs); i++ {
		if scs[i].Cid != sc.Cid {
			continue
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"msg":    "所选的课程,您已经选取",
			})
			return
		}
	}

	tid, err2 := GetTid(sc.Cid)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err2.Error(),
		})
		return
	}
	sc.Tid = tid
	db := mdb.Db
	err3 := db.C(models.CollectionStudent_Course).Insert(sc)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err3.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Insert Success!",
	})
}

func Avg(sid string) (avg int, err error) {
	db := mdb.Db
	var scs []models.Student_Course
	err = db.C(models.CollectionStudent_Course).Find(bson.M{"sid": sid}).All(scs)
	for i := 0; i < len(scs); i++ {
		avg += scs[i].Score
	}
	avg = avg / len(scs)
	return
}

//找出平均成绩排名前10的学生
//db.student_course.aggregate({$group:{_id:"$sid",nameAvg:{$avg:"$score"}}},{$sort: {nameAvg: -1}})
func GetForwardTen(c *gin.Context) {
	db := mdb.Db
	var s []interface{}
	m := []bson.M{
		{"$group": bson.M{"_id": "$sid", "nameAvg": bson.M{"$avg": "$score"}}},
		{"$sort": bson.M{"nameAvg": -1}},
		{"$limit": 10},
	}
	err := db.C(models.CollectionStudent_Course).Pipe(m).All(&s)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    s,
	})
	fmt.Println(456)
}

//找出选课数目排名前10的学生
func GetForwardTenByCount(c *gin.Context) {
	db := mdb.Db
	var s []interface{}
	m := []bson.M{
		{"$group": bson.M{"_id": "$sid", "count": bson.M{"$sum": 1}}},
		{"$sort": bson.M{"count": -1}},
		{"$limit": 10},
	}
	err := db.C(models.CollectionStudent_Course).Pipe(m).All(&s)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    s,
	})
}

//找出每位同学的最高成绩以及最高成绩对应的课程名
//db.student_course.aggregate([{$sort: {"sid":  -1, "score": -1}},{$group:{_id:"$sid",cid:{$first:"$cid"},score:{$first:"$score"}}}])
func GetMaxScore(c *gin.Context) {
	//i, _ := strconv.Atoi(c.Query("order"))
	//j, _ := strconv.Atoi(c.Query("limit"))
	//skip := (i - 1) * j
	db := mdb.Db
	var s []interface{}
	m := []bson.M{
		{"$sort": bson.M{"sid": -1, "score": -1}},
		{"$group": bson.M{"_id": "$sid", "cid": bson.M{"$first": "$cid"}, "score": bson.M{"$first": "$score"}}},
		{"$sort": bson.M{"_id": -1}},
	}
	err := db.C(models.CollectionStudent_Course).Pipe(m).All(&s)
	if err != nil {
		fmt.Println(err.Error())
	}
	//tem := s[skip : skip+j]
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    s,
	})
}

//求每门课程的选修人数和平均成绩
func GetCourseCountAvg(c *gin.Context) {
	db := mdb.Db
	var s []interface{}
	m := []bson.M{
		{"$group": bson.M{"_id": "$cid", "count": bson.M{"$sum": 1}, "avgScore": bson.M{"$avg": "$score"}}},
		{"$sort": bson.M{"count": -1}},
	}
	err := db.C(models.CollectionStudent_Course).Pipe(m).All(&s)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    s,
	})
}

//求每门课程最高成绩以及最高成绩对应的学生姓名
func GetCourseMaxStudent(c *gin.Context) {
	db := mdb.Db
	var s []interface{}
	m := []bson.M{
		{"$sort": bson.M{"cid": -1, "score": -1}},
		{"$group": bson.M{"_id": "$cid", "maxScore": bson.M{"$first": "$score"}, "sid": bson.M{"$first": "$sid"}}},
		{"$sort": bson.M{"_id": 1}},
	}
	err := db.C(models.CollectionStudent_Course).Pipe(m).All(&s)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    s,
	})
}

//求平均成绩排名前10的课程
func GetCourseByAvg(c *gin.Context) {
	db := mdb.Db
	var s []interface{}
	m := []bson.M{
		{"$group": bson.M{"_id": "$cid", "avgScore": bson.M{"$avg": "$score"}}},
		{"$sort": bson.M{"avgScore": -1}},
		{"$limit": 10},
	}
	err := db.C(models.CollectionStudent_Course).Pipe(m).All(&s)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    s,
	})
}

//求选课人数排名前10的课程
func GetCourseBySum(c *gin.Context) {
	db := mdb.Db
	var s []interface{}
	m := []bson.M{
		{"$group": bson.M{"_id": "$cid", "count": bson.M{"$sum": 1}}},
		{"$sort": bson.M{"count": -1}},
		{"$limit": 10},
	}
	err := db.C(models.CollectionStudent_Course).Pipe(m).All(&s)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    s,
	})
}

//列出student_course表中出现过的所有课程名称
func GetDistinctCourse(c *gin.Context) {
	i, _ := strconv.Atoi(c.Query("order"))
	j, _ := strconv.Atoi(c.Query("limit"))
	skip := (i - 1) * j

	db := mdb.Db
	var res []string
	err := db.C(models.CollectionStudent_Course).Find(nil).Distinct("cid", &res)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "error",
		})
		return
	}
	var courses []models.Course
	var template models.Course
	for i := skip; i < skip+j; i++ {
		err = db.C(models.CollectionCourse).Find(bson.M{"cid": res[i]}).One(&template)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "error",
			})
			return
		}
		courses = append(courses, template)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    courses,
		"count":   len(res),
	})
}
