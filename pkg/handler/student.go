package handler

import (
	mdb "MongodbPro/db"
	"MongodbPro/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

func GetStudentsTotal(c *gin.Context) {
	db := mdb.Db
	n, err := db.C(models.CollectionStudent).Count()
	fmt.Println("student表的条目数:", n)
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

func GetStudent(c *gin.Context) {
	i, _ := strconv.Atoi(c.Query("order"))
	j, _ := strconv.Atoi(c.Query("limit"))
	skip := (i - 1) * j
	db := mdb.Db
	var students []models.Student
	err := db.C(models.CollectionStudent).Find(bson.M{}).Skip(skip).Limit(j).All(&students)
	count, _ := db.C(models.CollectionStudent).Find(bson.M{}).Count()
	//fmt.Println(students[0].Birthday)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   students,
		"count":  count,
	})
}
func GetStudentByAge(c *gin.Context) {
	db := mdb.Db
	var students []models.Student
	fmt.Println("参数为:", c.Query("age"))
	age, err1 := strconv.Atoi(c.Query("age"))
	i, err2 := strconv.Atoi(c.Query("order"))
	j, err3 := strconv.Atoi(c.Query("limit"))
	skip := (i - 1) * j
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err1.Error(),
		})
		return
	}

	err := db.C(models.CollectionStudent).Find(bson.M{"age": bson.M{"$lte": age}}).Skip(skip).Limit(j).All(&students)
	count, _ := db.C(models.CollectionStudent).Find(bson.M{"age": bson.M{"$lte": age}}).Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   students,
		"count":  count,
	})

}
func GetStudentByAgeDname(c *gin.Context) {
	db := mdb.Db
	var students []models.Student
	i, _ := strconv.Atoi(c.Query("order"))
	j, _ := strconv.Atoi(c.Query("limit"))
	skip := (i - 1) * j

	age, err := strconv.Atoi(c.Query("age"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	query := bson.M{
		"age": bson.M{
			"$lte": age,
		},
		"dname": c.Query("dname"),
	}
	err1 := db.C(models.CollectionStudent).Find(query).Skip(skip).Limit(j).All(&students)
	count, _ := db.C(models.CollectionStudent).Find(query).Count()
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err1.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
		"data":   students,
		"count":  count,
	})
}

//得到所有学生的姓名年龄
func GetStudentNameAge(c *gin.Context) {
	db := mdb.Db
	var students []interface{}
	err := db.C(models.CollectionStudent).Find(bson.M{}).Select(bson.M{"name": 1, "age": 1, "_id": 0}).All(&students)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
		"data":   students,
	})
}

func GetNameSexByAge(c *gin.Context) {
	db := mdb.Db
	var students []interface{}
	age, err1 := strconv.Atoi(c.Query("age"))
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err1.Error(),
		})
		return
	}
	err := db.C(models.CollectionStudent).Find(bson.M{"age": bson.M{"$lte": age}}).Select(bson.M{"name": 1, "sex": 1, "_id": 0}).All(&students)
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
		"data":   students,
	})
}

func InsertStudent(c *gin.Context) {
	d, _ := c.GetRawData()
	log.Printf("Insert Student data: %s", string(d))
	st := models.Student{}
	_ = json.Unmarshal(d, &st)
	db := mdb.Db
	err := db.C(models.CollectionStudent).Insert(st)
	if err != nil {
		log.Panicf("Insert error: %s", err.Error())
	}
}

//前端的输入参数为要修改的一整个对象,即使只是修改某一个属性.也要传递一整个对象
func UpdateStudent(c *gin.Context) {
	db := mdb.Db
	dd, _ := c.GetRawData()
	var s models.Student
	err := json.Unmarshal(dd, &s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("年龄", s.Age)
	fmt.Println("更新学生的方法被调用!!")
	fmt.Println("生日:", s.Birthday)
	fmt.Println("sid:", s.Sid)
	selector := bson.M{"sid": s.Sid}
	data := bson.M{"$set": bson.M{"name": s.Name, "sex": s.Sex, "dname": s.Dname, "age": s.Age, "birthday": s.Birthday, "class": s.Class}}
	err1 := db.C(models.CollectionStudent).Update(selector, data)
	if err1 != nil {
		fmt.Println(123, err1.Error())
	}
}

func GetCourseInfo(c *gin.Context) {
	db := mdb.Db
	var course models.Course
	var courses []models.Course
	scs, err1 := GetCourseBysid(c.Query("sid"))
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err1.Error(),
		})
		return
	}
	fmt.Println(scs)
	fmt.Println("length", len(scs))
	for i := 0; i < len(scs); i++ {
		err2 := db.C(models.CollectionCourse).Find(bson.M{"cid": scs[i].Cid}).One(&course)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"msg":    err2.Error(),
			})
			return
		}
		courses = append(courses, course)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
		"data":   courses,
	})
}
func GetStudentBySid(c *gin.Context) {
	db := mdb.Db
	var student models.Student
	err := db.C(models.CollectionStudent).Find(bson.M{"sid": c.Query("sid")}).One(&student)
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
		"data":   student,
	})
}
func HandleFile(c *gin.Context) {
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
	fmt.Println("5e", xlsx.GetCellValue("Sheet1", "e5"))
	rows := xlsx.GetRows("Sheet1")
	for irow, row := range rows {
		if irow > 0 {
			var st models.Student
			st.Sid = row[0]
			st.Name = row[1]
			st.Sex = row[2]
			st.Age, _ = strconv.Atoi(row[3])
			//oc, _ := time.LoadLocation("Local")
			//st.Birthday, _ = time.ParseInLocation("01-02-2006", row[4], oc)
			st.Birthday = ConvertToFormatDay(row[4])
			st.Dname = row[5]
			st.Class = row[6]
			err3 := db.C(models.CollectionStudent).Insert(&st)
			if err3 != nil {
				fmt.Println(err3.Error())
				return
			}
		}
	}
	fmt.Println(header.Filename)
	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}
func ConvertToFormatDay(excelDaysString string) time.Time {
	// 2006-01-02 距离 1900-01-01的天数
	baseDiffDay := 38719 //在网上工具计算的天数需要加2天，什么原因没弄清楚
	curDiffDay := excelDaysString
	b, _ := strconv.Atoi(curDiffDay)
	// 获取excel的日期距离2006-01-02的天数
	realDiffDay := b - baseDiffDay
	//fmt.Println("realDiffDay:",realDiffDay)
	// 距离2006-01-02 秒数
	realDiffSecond := realDiffDay * 24 * 3600
	//fmt.Println("realDiffSecond:",realDiffSecond)
	// 2006-01-02 15:04:05距离1970-01-01 08:00:00的秒数 网上工具可查出
	baseOriginSecond := 1136185445
	resultTime := time.Unix(int64(baseOriginSecond+realDiffSecond), 0)
	return resultTime
}
