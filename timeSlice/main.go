package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	// 日志文件路径和命名规则
	filePath := "logs/"
	fileName := "app"
	logPattern := filePath + fileName + "-%Y%m%d-%H%M.log"
	// 这个包也行 但是用的人比较少
	//"github.com/pochard/logrotator"
	//writer, err := logrotator.NewTimeBasedRotator(logPattern, time.Minute*1)

	// 创建基于时间的日志轮转器
	writer, err := rotatelogs.New(
		logPattern, // 日志文件名模式
		rotatelogs.WithRotationTime(time.Minute*1), // 轮转时间间隔
		rotatelogs.WithMaxAge(time.Hour*24*7),      // 日志文件最大保留时间（7天）
	)
	if err != nil {
		panic(err)
	}

	// 配置 logrus 使用日志轮转器
	log := logrus.New()
	log.SetOutput(writer)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.Info("每分钟执行的任务触发")
	// 这里添加需要定时执行的业务逻辑
	log.WithFields(logrus.Fields{
		"event": "app_start",
		"user":  "admin",
	}).Info("Application started.")

	// 新增定时任务
	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			log.Info("每分钟执行的任务触发")
			// 这里添加需要定时执行的业务逻辑
			log.WithFields(logrus.Fields{
				"event": "app_start",
				"user":  "admin",
			}).Info("Application started.")
		}
	}()
	time.Sleep(time.Minute * 5)
}
