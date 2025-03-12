package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

// 这里实现的是将err级别的日志写入指定文件  其他的放在一起
// 但是error还是会放在一起  有待调整

type MyHook struct {
	Writer *os.File
}

func (hook *MyHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	hook.Writer.Write([]byte(line))
	return nil
}

func (hook *MyHook) Levels() []logrus.Level {
	// 将级别设置为Error级别
	return []logrus.Level{logrus.ErrorLevel}
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	logrus.SetReportCaller(true)
	file, _ := os.OpenFile("logs/err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	hook := &MyHook{Writer: file}
	logrus.AddHook(hook)
	logrus.Debugln("Debugln")
	logrus.Infoln("Infoln")
	logrus.Warnln("Warnln")
	logrus.Errorln("Errorln")
	logrus.Println("Println")
}

/*这里是修改 不在控制台输出err的代码
 */

/*


func main() {
    // 关闭默认输出
    logrus.SetOutput(io.Discard)

    // 添加控制台 Hook（输出非 ERROR 日志）
    logrus.AddHook(&ConsoleHook{
        Writer: os.Stdout,
        LogLevels: []logrus.Level{
            logrus.DebugLevel,
            logrus.InfoLevel,
            logrus.WarnLevel,
        },
    })

    // 添加文件 Hook（输出 ERROR 日志）
    file, _ := os.OpenFile("logs/err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    logrus.AddHook(&MyHook{Writer: file})

    // ... 其他测试代码保持不变 ...
}

// 新增控制台 Hook
type ConsoleHook struct {
    Writer    io.Writer
    LogLevels []logrus.Level
}

func (h *ConsoleHook) Levels() []logrus.Level {
    return h.LogLevels
}

func (h *ConsoleHook) Fire(entry *logrus.Entry) error {
    line, _ := entry.String()
    h.Writer.Write([]byte(line))
    return nil
}












*/
