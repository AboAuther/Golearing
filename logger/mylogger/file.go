package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	level       LogLevel
	filePath    string //文件保存的路径
	fileObj     *os.File
	errFileObj  *os.File
	fileName    string //文件名
	maxFileSize int64
}

//NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, fs int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: fs,
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}
func (f *FileLogger) initFile() error { //初始化文件， 连接文件名，打开两个日志文件。
	fullFileName := path.Join(f.filePath, f.fileName)
	fileobj, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,Error:%s", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,Error:%s", err)
		return err
	}
	f.fileObj = fileobj
	f.errFileObj = errFileObj
	f.Close()
	return nil
}

func (f *FileLogger) enable(loglevel LogLevel) bool { //类别比较
	return f.level <= loglevel
}

// 格式化输出到log日志文件中
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%v] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:03"), lv, fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			//如果要记录的日志级别大于等于Error
			//在ERROR文件中再记录一次
			fmt.Fprintf(f.errFileObj, "[%s] [%v] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:03"), lv, fileName, funcName, lineNo, msg)
		}
	}
}

//5类日志调用方法
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
