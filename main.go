package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func NewLog() *logrus.Logger {
	mLog := logrus.New()                       //新建一个实例
	mLog.SetOutput(os.Stdout)                  //设置输出类型
	mLog.SetReportCaller(true)                 //开启返回函数名和行号
	mLog.SetFormatter(&logrus.JSONFormatter{}) //设置自己定义的Formatter
	mLog.SetLevel(logrus.DebugLevel)           //设置最低的Level
	return mLog
}

func main() {
	log = NewLog()

	fmt.Println(log.GetLevel()) // info

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("hello world")
	log.Debugln("Debugln")
	log.Infoln("Infoln")
	log.Warnln("Warnln")
	log.Errorln("Errorln")
	log.Println("Println")
}

//   https://www.fengfengzhidao.com/article/jdlgH4sBEG4v2tWk8Wt7   作为参考 其中部分代码有点问题 注意辨别
