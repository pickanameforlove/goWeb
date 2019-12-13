package api

import "github.com/gin-gonic/gin"

func RunHTTPServer(engine *gin.Engine) {
	GetStudentAPI(engine)
	GetCourseAPI(engine)
	GetTeacherAPI(engine)
	GetStudentCourseAPI(engine)
}
