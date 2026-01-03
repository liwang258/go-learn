package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var File *os.File

func InitZapLogger(logPath string) *zap.Logger {
	initLogFile(logPath)
	// 1. 配置 zap
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //时间格式
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	//文件写入器
	fileCore := zapcore.NewCore(encoder, zapcore.AddSync(File), zapcore.DebugLevel)

	// 控制台写入器（可选）
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel)
	// 同时输出到文件和控制台
	core := zapcore.NewTee(fileCore, consoleCore)
	Logger = zap.New(core, zap.AddCaller())
	return Logger
}

func initLogFile(logPath string) {
	//1.创建日志文件
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	File = file
}
