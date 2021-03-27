package mylogger

import (
	"fmt"
	"time"
)

//往终端上写日志相关内容
type ConsoleLogger struct {
	Level LogLevel
}

//NewConsoleLogger 是构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(loglevel LogLevel) bool { //判断level大小，能否输出
	return c.Level <= loglevel
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) { //统一的log输出方法
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%v] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:03"), lv, fileName, funcName, lineNo, msg)
	}
}

//5类不同级别的日志调用方法
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
