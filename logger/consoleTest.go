package main

import (
	"./mylogger"
	"time"
)

func main() {
	log := mylogger.Newlog("Info")
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
