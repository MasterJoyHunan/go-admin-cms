package logger

import (
	"blog/pkg/setting"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

type MyHook struct{}

var openFile *os.File

func (h *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *MyHook) Fire(entry *logrus.Entry) (err error) {
	fileName := time.Now().Format("2006-01-02")
	fullPath := fmt.Sprintf("%s/%s.log", setting.LogConf.Path, fileName)

	// 无需多次获取文件句柄
	if openFile != nil && openFile.Name() == fullPath {
		return
	}

	if err = os.MkdirAll(setting.LogConf.Path, os.ModePerm); err != nil {
		log.Panic("创建文件夹错误", err)
		return
	}

	openFile, err = os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic("写入日志文件错误", err)
		return
	}

	//设置输出
	Logger.Out = openFile
	return
}
