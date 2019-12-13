package middleware

import (
	"MongodbPro/db"

	"github.com/gin-gonic/gin"
)

func Connect(context *gin.Context) {
	s := db.Session.Clone()
	defer s.Clone()

	context.Set("db", s.DB(db.Mongo.Database))

	context.Next()
}
