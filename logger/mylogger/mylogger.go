package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//自定义一个日志库
type LogLevel uint16

//Logger 接口
type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func NewLogger(model string) Logger {
	model = strings.ToLower(model)
	switch model {
	case "c":
		var input string
		fmt.Println("请输入日志级别：")
		fmt.Scanf("%s\n", &input)
		return NewConsoleLogger(input)
	case "f":
		var level, ft, fn string
		var fs int64
		fmt.Println("请输入日志级别：")
		fmt.Scanf("%s\n", &level)
		fmt.Println("请输入日志文件路径：")
		fmt.Scanf("%s\n", &ft)
		fmt.Println("请输入日志文件名：")
		fmt.Scanf("%s\n", &fn)
		fmt.Println("请输入日志文件最大容量：")
		fmt.Scanf("%d\n", &fs)
		return NewFileLogger(level, ft, fn, fs)
	default:
		return nil
	}

}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("未知日志级别")
		return UNKNOWN, err
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}
