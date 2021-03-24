package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("file open failed,Error:%v", err)
		return
	}

}
