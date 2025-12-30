package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Handle(c *gin.Context) {
	fmt.Println("收到用户相关请求")
}
