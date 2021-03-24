package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("File open failed,Error:%v", err)
		return
	}

	//关闭文件
	defer fileObj.Close()
	for {
		//读文件
		//var tmp =make([]byte,128)//指定读的长度
		var tmp [128]byte
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Printf("读取完毕 ,Error:%v", err)
			return
		}
		if err != nil {
			fmt.Printf("File read failed,Error:%v", err)
			return
		}
		//fmt.Printf("读了%d个字节",n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}

}

//利用bufio 这个包读文件
func readFromFileByBufio() {

	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("File open failed,Error:%v", err)
		return
	}

	//关闭文件
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		//创建一个用来从文件中读内容的对象

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("File read failed,Error:%v", err)
			return
		}
		fmt.Print(line)
	}
}

//利用ioutil 包读文件
func readFromFileByIoutil() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("File open failed,Error:%v", err)
		return
	}

	//关闭文件
	defer fileObj.Close()
	res, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file faild,Eroor:%v", err)
		return
	}
	fmt.Println(string(res))

}

//打开文件
func main() {

	//readFromFile()
	//readFromFileByBufio()
	readFromFileByIoutil()

}
