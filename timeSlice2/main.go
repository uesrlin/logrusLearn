package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

// 采用hook的方式  实现两个接口

type FileDateHook struct {
	mu sync.RWMutex // 新增读写锁
	// 日志文件
	file *os.File
	// 日志文件路径
	path string
	// 时间
	date string
	// 应用名称
	appName string
}

func (hook *FileDateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (hook *FileDateHook) Fire(entry *logrus.Entry) error {

	line, _ := entry.String()
	// 这个 entry.Time 是 logrus 在创建 Entry 时自动设置的当前时间
	currentTime := entry.Time.Format("2006-01-02_15-04")

	// 使用读写锁保证线程安全
	hook.mu.Lock()
	defer hook.mu.Unlock()

	if hook.date == currentTime {
		if _, err := hook.file.Write([]byte(line)); err != nil {
			logrus.Errorf("写入日志失败: %v", err)
		}
		return nil
	}

	// 关闭旧文件
	if hook.file != nil {
		if err := hook.file.Close(); err != nil {
			logrus.Errorf("关闭日志文件失败: %v", err)
		}
	}

	// 打开新文件（添加错误重试机制）
	newFile, err := os.OpenFile(getFilePath(hook.path, currentTime, hook.appName),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Errorf("打开日志文件失败: %v", err)
		return err
	}

	// 更新文件状态
	hook.file = newFile
	hook.date = currentTime

	// 写入当前日志
	if _, err := hook.file.Write([]byte(line)); err != nil {
		logrus.Errorf("写入日志失败: %v", err)
	}
	return nil

}
func InitFile(logPath string, appName string) error {
	// 参数校验
	if logPath == "" || appName == "" {
		return fmt.Errorf("日志路径和应用名称不能为空")
	}

	// 使用当前时间初始化
	initialTime := time.Now().Format("2006-01-02_15-04")
	err2 := os.MkdirAll(logPath, 0755)
	if err2 != nil {
		return fmt.Errorf("创建日志目录失败: %v", err2)
	}

	// 直接调用公共路径生成方法
	filePath := getFilePath(logPath, initialTime, appName)
	fmt.Println(filePath)

	// 创建文件（添加错误处理）
	file, err := os.OpenFile(filePath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("初始化日志文件失败: %v", err)
	}

	// 初始化hook
	logrus.AddHook(&FileDateHook{
		file:    file,
		path:    logPath,
		date:    initialTime,
		appName: appName,
	})
	return nil
}

// 提取公共文件路径生成逻辑
func getFilePath(logPath, timestamp, appName string) string {
	return fmt.Sprintf("%s/%s-%s.log", logPath, timestamp, appName)
}
func main() {
	err := InitFile("log", "myapp")
	if err != nil {
		logrus.Errorf("初始化日志失败: %v", err)
	}
	// 测试日志写入
	for {
		logrus.Errorf("示例错误日志")
		logrus.Warn("警告日志")
		time.Sleep(20 * time.Second)
	}

}
