package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main()  {
	app := &cli.App{
		Name: "arguments",
		Usage: "arguments example",
		Action: func(c *cli.Context) error {
			for i := 0; i < c.NArg(); i++ {
				fmt.Printf("%d: %s\n",i+1,c.Args().Get(i)) //%d 表示为十进制 %s 直接输出字符串或者[]byte
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
