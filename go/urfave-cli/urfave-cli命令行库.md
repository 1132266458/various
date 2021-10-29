初始化mod文件

```shell
go get cli
```



安装cli库

```shell
go get -u github.com/urfave/cli/v2
```



代码

````go
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
````





编译、运行

```shell
go build -o hello
运行./hello
输出
hello world

运行./hello --help 
输出
NAME:
   hello - hello world example

USAGE:
   hello [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```



**参数**

通过`cli.Context`的相关方法我们可以获取传给命令行的参数信息：

- `NArg()`：返回参数个数；
- `Args()`：返回`cli.Args`对象，调用其`Get(i)`获取位置`i`上的参数。

示例：

```go
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
```

这里只是简单输出：

```shell
$ go run main.go hello world
1: hello
2: world
```



**选项**
一个好用的命令行程序怎么会少了选项呢？`cli`设置和获取选项非常简单。在`cli.App{}`结构初始化时，设置字段`Flags`即可添加选项。`Flags`字段是`[]cli.Flag`类型，`cli.Flag`实际上是接口类型。`cli`为常见类型都实现了对应的`XxxFlag`，如`BoolFlag/DurationFlag/StringFlag`等。它们有一些共用的字段，`Name/Value/Usage`（名称/默认值/释义）。看示例：

```go
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
```



**存入变量**

除了通过`c.Type(name)`来获取选项的值，我们还可以将选项存到某个预先定义好的变量中。只需要设置`Destination`字段为变量的地址即可：

```go
package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var language string

func main()  {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "lang",
				Value: "english",
				Usage: "language for the greeting",
				Destination: &language,
			},
		},
		Action: func(c *cli.Context) error {
			name := "world"
			if c.NArg() > 0{
				name = c.Args().Get(0)
			}
			if language == "english" {
				fmt.Println("helolo",name)
			}else{
				fmt.Println("你好",name)
			}

			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
```

