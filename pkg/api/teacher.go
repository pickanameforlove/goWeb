package api

import (
	"MongodbPro/pkg/handler"

	"github.com/gin-gonic/gin"
)

func GetTeacherAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1alpha")
	v1.GET("teacher/", handler.GetTeacher)
	v1.GET("/teacher/count", handler.GetTeachersTotal)
	v1.GET("teacher/getTeacherBySex", handler.GetTeacherBySex)
	v1.GET("teacher/getTeacherByDname", handler.GetTeacherByDname)
	v1.POST("teacher/update", handler.UpdateTeacher)
	v1.GET("/teacher/getTeacherByAge", handler.GetTeacherByAge)
	v1.POST("/teacher/insert", handler.InsertTeacher)
	v1.POST("/teacher/file", handler.HandleTeacherFile)
}
