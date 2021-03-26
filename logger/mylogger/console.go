package mylogger

import (
	"fmt"
	"time"
)

//往终端上写日志相关内容
type Logger struct {
	Level LogLevel
}

func Newlog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(loglevel LogLevel) bool {
	return l.Level <= loglevel
}

func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%v] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:03"), lv, fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}

}
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}

}
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		log(WARNING, msg)
	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		log(FATAL, msg)
	}
}
