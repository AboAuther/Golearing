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
	return nil
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file information failed,Error:%s", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) enable(loglevel LogLevel) bool { //类别比较
	return f.level <= loglevel
}

func (f *FileLogger) spiltFile(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file information failed,Error:%s", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	//需要切割文件
	//1.关闭当前的日志文件
	file.Close()
	//2.备份一下 rename  xx.log->xx.log.bak20210327
	os.Rename(logName, newLogName)
	//3.打开一个新文件
	fileObj, err := os.OpenFile(f.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,Error:%v", err)
		return nil, err
	}
	//4.将新日志文件对象赋值给 f.fileobj
	return fileObj, nil
}

// 格式化输出到log日志文件中
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFile, err := f.spiltFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%v] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:03"), lv, fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			//如果要记录的日志级别大于等于Error
			//在ERROR文件中再记录一次
			if f.checkSize(f.errFileObj) {
				newFile, err := f.spiltFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}

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
