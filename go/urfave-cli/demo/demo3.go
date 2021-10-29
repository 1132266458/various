package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main()  {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
			&cli.StringFlag{
				Name: "color",
				Value: "red",
				Usage: "color for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			name := "world"
			if c.NArg() > 0{
				name = c.Args().Get(0)
			}
			if c.String("lang") == "english" {
				fmt.Println("helolo",name)
			}else{
				fmt.Println("你好",name)
			}

			if c.String("color") == "red" {
				fmt.Println("input red")
			}else{
				fmt.Println(c.String("color"))
			}

			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
