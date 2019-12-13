package handler

import (
	mdb "MongodbPro/db"
	"MongodbPro/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func GetCoursesTotal(c *gin.Context) {
	db := mdb.Db
	n, err := db.C(models.CollectionCourse).Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "success",
		"data":   n,
	})
}

func GetCourse(c *gin.Context) {
	i, _ := strconv.Atoi(c.Query("order"))
	j, _ := strconv.Atoi(c.Query("limit"))
	skip := (i - 1) * j
	db := mdb.Db
	var courses []models.Course
	count, _ := db.C(models.CollectionCourse).Find(bson.M{}).Count()
	err := db.C(models.CollectionCourse).Find(bson.M{}).Skip(skip).Limit(j).All(&courses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
		"data":   courses,
		"count":  count,
	})
}
func GetCourseByFcid(c *gin.Context) {
	i, _ := strconv.Atoi(c.Query("order"))
	j, _ := strconv.Atoi(c.Query("limit"))
	skip := (i - 1) * j
	db := mdb.Db
	var courses []models.Course
	count, _ := db.C(models.CollectionCourse).Find(bson.M{"fcid": c.Query("fcid")}).Count()
	err := db.C(models.CollectionCourse).Find(bson.M{"fcid": c.Query("fcid")}).Skip(skip).Limit(j).All(&courses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
		"data":   courses,
		"count":  count,
	})
}

//课程对象插入.
func InsertCourse(c *gin.Context) {
	d, _ := c.GetRawData()
	log.Printf("Insert Course data: %s", string(d))
	cs := models.Course{}
	_ = json.Unmarshal(d, &cs)
	db := mdb.Db
	err := db.C(models.CollectionCourse).Insert(cs)
	if err != nil {
		log.Panicf("Insert error: %s", err.Error())
	}
}

//前端的输入参数为要修改的一整个对象,即使只是修改某一个属性.也要传递一整个对象
func UpdateCourse(c *gin.Context) {
	db := mdb.Db
	dd, _ := c.GetRawData()
	var s models.Course
	err := json.Unmarshal(dd, &s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.C(models.CollectionCourse).Update(bson.M{"cid": s.Cid}, bson.M{"$set": bson.M{"name": s.Name, "credit": s.Credit, "fcid": s.Fcid}})
}
func GetCourseNonHttp(cid string) (course models.Course, err error) {
	db := mdb.Db
	err = db.C(models.CollectionCourse).Find(bson.M{"cid": cid}).One(&course)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
	return
}
func HandleCourseFile(c *gin.Context) {
	db := mdb.Db
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败"})
		return
	}
	xlsx, err1 := excelize.OpenReader(file)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
		return
	}
	rows := xlsx.GetRows("Sheet1")
	for irow, row := range rows {
		if irow > 0 {
			var course models.Course
			course.Cid = row[0]
			course.Name = row[1]
			course.Fcid = row[2]
			course.Credit = row[3]

			err3 := db.C(models.CollectionCourse).Insert(&course)
			if err3 != nil {
				fmt.Println(err3.Error())
				return
			}
		}
	}
	fmt.Println(header.Filename)
	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}
