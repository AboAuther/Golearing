package main

import (
	"fmt"
	"io"
	"os"
)

func copyfile(srcName, desName string) (written int64, err error) {
	//只读方式放开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("file %s open failed,Error:%d", srcName, err)
		return
	}
	defer src.Close()
	//以写|创建方式打开目标文件
	des, err := os.OpenFile(desName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("file %s open failed,Error:%d", desName, err)
		return
	}
	defer des.Close()
	return io.Copy(des, src)

}

func main() {
	_, err := copyfile("xx.txt", "des.txt")
	if err != nil {
		fmt.Printf("file copy failed,Error:%d", err)
		return
	}
	fmt.Println("copy done!")

}
