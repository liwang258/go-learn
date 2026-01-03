package biz

import (
	"go-learn/task4"
	"go-learn/task4/biz/user"
	"net/http"
	"strings"

	mylogger "go-learn/task4/biz/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func RegisterRouter() *gin.Engine {
	//将 Gin 日志重定向到 zap
	gin.DefaultWriter = zapcore.AddSync(mylogger.File)
	gin.DefaultErrorWriter = zapcore.AddSync(mylogger.File)
	// 设置 Gin 模式（生产环境建议 release）
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(TraceIdMiddleware())
	// 注册路由
	router.Group("user").POST("/create", func(c *gin.Context) {
		var query task4.UserQuery
		//c.Bind(&struct)：与 ShouldBind 类似，但绑定失败时会直接返回 400 Bad Request，不进入处理函数。
		//c.ShouldBindQuery(&struct)：专门绑定 URL 查询参数。
		//c.ShouldBindJSON(&struct)：专门绑定 JSON 请求体。
		//c.ShouldBind(&struct)：自动根据请求类型（GET、POST 等）选择绑定方式（form、json 等）。
		//自动绑定请求，如果绑定失败，说明参数非法，这里直接就返回400错误
		if !BindJSON(c, &query) {
			return
		}
		mylogger.Logger.Info("accept create user request", zap.Any("query", query))
		if result, err := user.CreateUser(query); err != nil {
			mylogger.Logger.Error("create user failed err", zap.Error(err), zap.Any("query", query))
			c.JSON(http.StatusBadRequest, task4.Response{
				Code: 400,
				Msg:  err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, task4.Response{
				Code: 200,
				Msg:  "Success",
				Data: result,
			})
		}

	})

	return router
}

func TraceIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//如果请求中有传traceId，则用请求传过来的，没有则重新生成一个
		traceId := c.Request.Header.Get("X-Trace-Id")
		if traceId == "" || strings.Trim(traceId, "\"") == "" {
			traceId = uuid.NewString()
		}
		//将traceId写入响应头(方便客户端查看)
		c.Writer.Header().Set("X-Trace-Id", traceId)
		//将traceId 存入gin上下文
		c.Set("traceId", traceId)
		reqLogger := mylogger.Logger.With(zap.String("traceId", traceId))
		//将子logger存入上下文，方便后续使用
		c.Set("logger", reqLogger)
		//记录请求开始
		reqLogger.Info("request started", zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path))
		//处理请求
		c.Next()

		reqLogger.Info("request completed", zap.Int("status", c.Writer.Status()))
	}

}

// BindJSON 统一 JSON 参数绑定与错误处理
func BindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		// 统一日志记录
		zap.L().Error("参数绑定失败",
			zap.Error(err),
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
		)

		// 统一错误响应
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数格式错误",
		})
		return false
	}
	return true
}
