package db

import (
	"go-learn/task4/config"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	instanc *dbInstance
	once    sync.Once
)

type dbInstance struct {
	DB *gorm.DB
}

func InitDb(config *config.Config) {
	once.Do(func() {
		// 1. 配置 GORM 日志：打印最终执行的 SQL 到控制台
		newLogger := logger.New(
			// 将日志输出到控制台（os.Stdout）
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值（超过 1 秒会标红）
				LogLevel:                  logger.Info, // 日志级别：Info 会打印所有 SQL
				IgnoreRecordNotFoundError: true,        // 忽略记录不存在的错误
				Colorful:                  true,        // 开启彩色输出（控制台更易读）
			},
		)
		dsn := config.Database.DSN
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			panic(err)
		}
		instanc = &dbInstance{
			DB: db,
		}
	})
}

func GetDb() *gorm.DB {
	return instanc.DB
}
