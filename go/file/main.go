package main

import (
	"fmt"
	"os"
)

func main()  {
	path := "./various/go/file/test"
	err := os.MkdirAll(path,0666) //创建多级
	if err != nil {
		panic(err)
	}

	f,err := os.Create(path+"/test.txt")
	defer f.Close()
	if err !=nil {
		fmt.Println(err.Error())
	} else {
		_,err=f.Write([]byte("要写入的文本内容"))
	}
}
