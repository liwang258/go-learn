package main

import (
	"context"
	"flag"
	"go-learn/task4/biz"
	"go-learn/task4/config"
	"go-learn/task4/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/trace"
	"syscall"
	"time"
)

func init() {
	path := flag.String("config", "script/config.yml", "配置文件路径默认 script/config.yml")
	if cfg, err := config.LoadConfig(*path); err != nil {
		panic(err)
	} else {
		//初始化数据库
		db.InitDb(cfg)
	}
}

func main() {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
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
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	trace.Log(context.Background(), "[start]", "server start ok")

	// 等待退出信号（优雅关闭）
	quit := make(chan os.Signal, 1)
	// 监听 SIGINT 和 SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 创建一个 5 秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭 HTTP 服务
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")
}
