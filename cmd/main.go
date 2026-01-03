package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"go-learn/task4/biz"
	"go-learn/task4/biz/logger"
	"go-learn/task4/config"
	"go-learn/task4/db"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func init() {
	path := flag.String("config", "script/config.yml", "配置文件路径默认 script/config.yml")
	cfg, err := config.LoadConfig(*path)
	if err != nil {
		panic(err)
	}

	//初始化数据库
	db.InitDb(cfg)
}

func main() {

	//1.创建日志文件
	logger.InitZapLogger("trace.out")

	defer func(logger *zap.Logger) {
		if err := logger.Sync(); err != nil {
			fmt.Println("sync log file failed!", err)
		}
	}(logger.Logger) // 确保日志刷新到磁盘

	defer func(File *os.File) {
		err := File.Close()
		if err != nil {
			panic(err)
		}
	}(logger.File)

	//注册路由策略
	router := biz.RegisterRouter()

	//启动http服务，暴露端口为8080，指定http请求处理器为 router.Handler
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}
	// 启动 HTTP 服务（异步）
	go func() {
		// 监听网络请求
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Fatal("服务启动失败: ", zap.Error(err))
		}
	}()
	logger.Logger.Info("服务启动成功 ")
	// 等待退出信号（优雅关闭）
	quit := make(chan os.Signal, 1)
	// 监听 SIGINT 和 SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("正在关闭服务器...")

	// 创建一个 5 秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭 HTTP 服务
	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Fatal("服务器强制关闭", zap.Error(err))
	}

	logger.Logger.Info("服务器已退出")
}
