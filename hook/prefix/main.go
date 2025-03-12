package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type MyHook struct {
	Prefix string
}

func (h *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (h *MyHook) Fire(entry *logrus.Entry) error {
	// 在这里处理日志记录
	//entry.Data的作用是为日志条目添加额外的上下文信息，这些信息可以在Formatter中被格式化输出
	entry.Data["app"] = "myapp"
	// 在消息前添加固定前缀
	entry.Message = fmt.Sprintf("[%s] %s", h.Prefix, entry.Message)

	return nil
}
func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	logrus.SetOutput(os.Stdout)
	logrus.AddHook(&MyHook{Prefix: "MYAPP"})
	logrus.Debugln("Debugln")
	logrus.Infoln("Infoln")
	logrus.Warnln("Warnln")
	logrus.Errorln("Errorln")
	logrus.Println("Println")
}

// 但是最后的输出是这样的，不太美观  输出顺序有待调整
/*
INFO[2025-03-12 19:11:07] [MYAPP] Infoln                                app=myapp
WARN[2025-03-12 19:11:07] [MYAPP] Warnln                                app=myapp
ERRO[2025-03-12 19:11:07] [MYAPP] Errorln                               app=myapp
INFO[2025-03-12 19:11:07] [MYAPP] Println                               app=myapp


*/
