package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.只读
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.只写
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.读写
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.追加
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.没有文件则创建
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.可以测试文件是否存在
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.每次打开都清空
*/
func writeDemo() {
	fileobj, err := os.OpenFile("./xx.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("file open failed,Error:%v", err)
		return
	}
	//write
	fileobj.Write([]byte("It's a beautiful day\n"))
	//writestring
	fileobj.WriteString("It's a rainy day\n")
	fileobj.Close()
}
func writeDemo2() {
	fileobj, err := os.OpenFile("./xx.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("file open failed,Error:%v", err)
		return
	}
	//创建一个写的对象
	wr := bufio.NewWriter(fileobj)
	wr.WriteString("It is a sunny day") //写到缓存
	wr.Flush()                          //将缓存中的内容写到文件中
}

func writeDemo3() {
	str := "will you marry me ?"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("file open failed,Error:%v", err)
		return
	}
}
func main() {
	//writeDemo2()
	//writeDemo()
	writeDemo3()

}
