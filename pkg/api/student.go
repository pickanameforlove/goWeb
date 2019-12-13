package api

import (
	"MongodbPro/pkg/handler"

	"github.com/gin-gonic/gin"
)

func GetStudentAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1alpha")
	//v1.Use(middleware.JWTAuth())
	v1.GET("/student/count", handler.GetStudentsTotal)
	v1.GET("student/", handler.GetStudent)
	v1.GET("student/byAge", handler.GetStudentByAge)
	v1.GET("student/byAgeDname", handler.GetStudentByAgeDname)
	v1.GET("student/getNameAge", handler.GetStudentNameAge)
	v1.GET("student/getNameSexByAge", handler.GetNameSexByAge)
	v1.POST("/student/update", handler.UpdateStudent)
	v1.GET("/student/courses", handler.GetCourseInfo)
	v1.POST("/student/insert", handler.InsertStudent)
	v1.POST("/file", handler.UploadHandler)
	v1.GET("/student/bySid", handler.GetStudentBySid)
	v1.POST("/student/file", handler.HandleFile)

}
