package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	path := "logs/"
	file, _ := os.OpenFile(path+"info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("hello world")
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Debug("hello debug")
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Infof("-----%s", "你好")
	//Infoln是用于打印信息级别的日志，参数会被当作空格分隔的值处理
	logrus.Infoln("nihao", "我不好", "我太不好了")

}
