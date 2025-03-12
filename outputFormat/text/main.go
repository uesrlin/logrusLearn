package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(logrus.GetLevel()) // info

	logrus.SetLevel(logrus.DebugLevel)
	// 这样是文本格式
	/*
		ForceColors：是否强制使用颜色输出。
		DisableColors：是否禁用颜色输出。
		ForceQuote：是否强制引用所有值。
		DisableQuote：是否禁用引用所有值。
		DisableTimestamp：是否禁用时间戳记录。
		FullTimestamp：是否在连接到 TTY 时输出完整的时间戳。
		TimestampFormat：用于输出完整时间戳的时间戳格式。
	*/
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("hello world")
	logrus.Debugln("Debugln")
	logrus.Infoln("Infoln")
	logrus.Warnln("Warnln")
	logrus.Errorln("Errorln")
	logrus.Println("Println")

}
