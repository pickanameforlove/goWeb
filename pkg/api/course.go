package api

import (
	"MongodbPro/pkg/handler"

	"github.com/gin-gonic/gin"
)

func GetCourseAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1alpha")
	v1.GET("/course/count", handler.GetCoursesTotal)
	v1.GET("course/", handler.GetCourse)
	v1.GET("/course/getCourseByFcid", handler.GetCourseByFcid)
	v1.POST("/course/insert", handler.InsertCourse)
	v1.POST("/course/update", handler.UpdateCourse)
	v1.POST("/course/file", handler.HandleCourseFile)
}
