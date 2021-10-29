package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main()  {
	//创建一个cli.App结构的对象，然后调用其Run()方法，传入命令行的参数即可
	//Name和Usage都显示在帮助中，Action是调用该命令行程序时实际执行的函数
	app := &cli.App{
		Name: "hello",
		Usage: "hello world example",
		Action: func(c *cli.Context) error {
			fmt.Println("hello world")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)  //打印输出内容  退出应用程序  defer函数不会执行
	}
}
