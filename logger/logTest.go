package main

import (
	"./mylogger"
	"fmt"
	"time"
)

var log mylogger.Logger

func main() {
	/*log = mylogger.NewConsoleLogger("Info")
	log = mylogger.NewFileLogger("Info", "./", "Mylog.log", 10*1024*1024)
	*/

	var input string
	fmt.Println("请输入你要操作的模式（c:控制台日志，f:文件日志）：")
	fmt.Scanf("%s\n", &input)
	log = mylogger.NewLogger(input)

	for {
		name := "AboAuther"
		id := 10001
		log.Debug("这是一条debug日志")
		log.Info("这是一条Info日志,name:%s,id:%d", name, id)
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 3)
	}

}
