package biz

import (
	"go-learn/task4/biz/user"

	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	router := gin.Default()

	router.Group("user").GET("/", func(c *gin.Context) {
		user.Handle(c)
	})

	return router
}
