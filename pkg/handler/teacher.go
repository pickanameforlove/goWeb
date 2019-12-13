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

func GetTeachersTotal(c *gin.Context) {
	db := mdb.Db
	n, err := db.C(models.CollectionTeacher).Count()
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

func GetTeacher(c *gin.Context) {
	i, _ := strconv.Atoi(c.Query("order"))
	j, _ := strconv.Atoi(c.Query("limit"))
	skip := (i - 1) * j
	db := mdb.Db
	var teachers []models.Teacher
	count, _ := db.C(models.CollectionTeacher).Find(bson.M{}).Count()
	err := db.C(models.CollectionTeacher).Find(bson.M{}).Skip(skip).Limit(j).All(&teachers)
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
		"data":   teachers,
		"count":  count,
	})
}

func GetTeacherBySex(c *gin.Context) {
	i, _ := strconv.Atoi(c.Query("order"))
	j, _ := strconv.Atoi(c.Query("limit"))
	skip := (i - 1) * j
	db := mdb.Db
	var teachers []models.Teacher
	fmt.Println(c.Query("sex"))
	count, _ := db.C(models.CollectionTeacher).Find(bson.M{"sex": c.Query("sex")}).Count()
	err := db.C(models.CollectionTeacher).Find(bson.M{"sex": c.Query("sex")}).Skip(skip).Limit(j).All(&teachers)
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
		"data":   teachers,
		"count":  count,
	})
}

func GetTeacherByDname(c *gin.Context) {
	i, _ := strconv.Atoi(c.Query("order"))
	j, _ := strconv.Atoi(c.Query("limit"))
	skip := (i - 1) * j
	db := mdb.Db
	var teachers []models.Teacher
	count, _ := db.C(models.CollectionTeacher).Find(bson.M{"dname": c.Query("dname")}).Count()
	err := db.C(models.CollectionTeacher).Find(bson.M{"dname": c.Query("dname")}).Skip(skip).Limit(j).All(&teachers)
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
		"data":   teachers,
		"count":  count,
	})
}
func GetTeacherByAge(c *gin.Context) {
	i, _ := strconv.Atoi(c.Query("order"))
	j, _ := strconv.Atoi(c.Query("limit"))
	skip := (i - 1) * j
	db := mdb.Db
	var teachers []models.Teacher
	age, _ := strconv.Atoi(c.Query("age"))
	count, _ := db.C(models.CollectionTeacher).Find(bson.M{"age": bson.M{"$gte": age}}).Count()
	err := db.C(models.CollectionTeacher).Find(bson.M{"age": bson.M{"$gte": age}}).Skip(skip).Limit(j).All(&teachers)
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
		"data":   teachers,
		"count":  count,
	})
}
func InsertTeacher(c *gin.Context) {
	d, _ := c.GetRawData()
	log.Printf("Insert Teacher data: %s", d)
	te := models.Teacher{}
	_ = json.Unmarshal(d, &te)
	db := mdb.Db
	err := db.C(models.CollectionTeacher).Insert(te)
	if err != nil {
		log.Panicf("Insert error: %s", err.Error())
	}
}

//前端的输入参数为要修改的一整个对象,即使只是修改某一个属性.也要传递一整个对象
func UpdateTeacher(c *gin.Context) {
	db := mdb.Db
	dd, _ := c.GetRawData()
	var s models.Teacher
	err := json.Unmarshal(dd, &s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err1 := db.C(models.CollectionTeacher).Update(bson.M{"tid": s.Tid}, bson.M{"$set": bson.M{"name": s.Name, "sex": s.Sex, "dname": s.Dname, "age": s.Age}})
	if err1 != nil {
		fmt.Println(err1.Error())
	}
}

//func findAge(str string) (age int) {
//	for i := 0; i < len(str); i++ {
//
//	}
//}
func HandleTeacherFile(c *gin.Context) {
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
			var teacher models.Teacher
			teacher.Tid = row[0]
			teacher.Name = row[1]
			teacher.Sex = row[2]
			teacher.Age, _ = strconv.Atoi(row[3])
			teacher.Dname = row[4]

			err3 := db.C(models.CollectionTeacher).Insert(&teacher)
			if err3 != nil {
				fmt.Println(err3.Error())
				return
			}
		}
	}
	fmt.Println(header.Filename)
	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}
