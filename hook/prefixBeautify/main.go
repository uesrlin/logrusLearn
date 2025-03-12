package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

//  暂时没用因为 这里重写了 textFormater
//type MyHook struct {
//}
//
//func (h *MyHook) Levels() []logrus.Level {
//	return logrus.AllLevels
//}
//func (h *MyHook) Fire(entry *logrus.Entry) error {
//	// 在这里处理日志记录
//	//entry.Data的作用是为日志条目添加额外的上下文信息，这些信息可以在Formatter中被格式化输出
//	entry.Data["app"] = "myapp"
//	return nil
//}

type CustomFormatter struct {
	logrus.TextFormatter
	Prefix string
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 创建缓冲区
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 设置颜色
	var levelColor int
	switch entry.Level {
	case logrus.WarnLevel:
		levelColor = 33 // 黄色
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // 红色
	case logrus.DebugLevel:
		levelColor = 36 // 青色
	default:
		levelColor = 34 // 蓝色
	}
	// 构建基础日志信息
	fmt.Fprintf(b, "[%s]", f.Prefix)
	fmt.Fprintf(b, "  \x1b[%dm[%s]\x1b[0m", levelColor, entry.Time.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(b, "  \x1b[%dm%-7s\x1b[0m", levelColor, strings.ToUpper(entry.Level.String())) // 带颜色的级别

	// 添加消息主体
	if entry.Message != "" {
		fmt.Fprintf(b, " - %s", entry.Message)
	}

	// 处理附加字段
	if len(entry.Data) > 0 {
		b.WriteString(" | ")
		fields := make([]string, 0, len(entry.Data))
		for k, v := range entry.Data {
			fields = append(fields, fmt.Sprintf("%s=%v", k, v))
		}
		b.WriteString(strings.Join(fields, " "))
	}

	// 处理调用者信息
	if entry.HasCaller() {
		fmt.Fprintf(b, " \x1b[90m(%s:%d)\x1b[0m",
			path.Base(entry.Caller.File), // 仅显示文件名
			entry.Caller.Line)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func main() {
	logrus.SetFormatter(&CustomFormatter{
		Prefix: "MYAPP",
		TextFormatter: logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	})
	// 在日志文件中是不会有颜色的所以不推荐  但是可以在终端中看到颜色  但是不确定采用submit的方式是否会有颜色
	//path := "logs/"
	//file, _ := os.OpenFile(path+"info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//logrus.SetOutput(file)

	logrus.SetOutput(os.Stdout)
	logrus.Debugln("Debugln")
	logrus.Infoln("Infoln")
	logrus.Warnln("Warnln")
	logrus.Errorln("Errorln")
	logrus.Println("Println")
}
