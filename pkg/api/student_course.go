package api

import (
	"MongodbPro/pkg/handler"

	"github.com/gin-gonic/gin"
)

func GetStudentCourseAPI(engine *gin.Engine) {
	v1 := engine.Group("/api/v1alpha")
	v1.POST("/sc/insert", handler.InsertSC)
	v1.GET("/test", handler.GetForwardTen)
	v1.GET("/test1", handler.GetForwardTenByCount)
	v1.GET("/test2", handler.GetMaxScore)
	v1.GET("/test3", handler.GetCourseCountAvg)
	v1.GET("/test4", handler.GetCourseMaxStudent)
	v1.GET("/test5", handler.GetCourseByAvg)
	v1.GET("/test6", handler.GetCourseBySum)
	v1.GET("/test7", handler.GetDistinctCourse)
}
