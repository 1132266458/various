# go build go install

**go build**

go build 入口是一个main包，有main包才能生产exe文件，一个mian包里只能有一个唯一的main方法。

go build产生的exe文件是在当前的项目主目录生成。

```html
go build -o rename.exe   //可以用-o参数来指定编译后可执行文件的名字
```



**go install**

go install 的作用有两步：第一步是编译导入的包文件，所有导入的包文件编译完才会编译主程序；第二步是将编译后生成的可执行文件放到 bin 目录下（$GOPATH/bin），编译后的包文件放到 pkg 目录（$GOPATH/pkg）。



# 第三方安装包的方法

**go get**

```html
比如要安装 "github.com/gin-gonic/gin"
go get github.com/gin-gonic/gin

注意：执行go get 命令需要先安装git命令，并配置git全局变量。
```



**源码包安装**

由于国内网络问题，很多时候go get命令并不能安装，所以就需要手动下载源码包，然后拷贝到$GOPATH/sr/ 目录下

```
比如要安装"github.com/julienschmidt/httprouter"
去github.com/julienschmidt下载源码包
拷贝到 $GOPATH/src/github.com/julienschmidt

cd $GOPATH/src/github.com/julienschmidt/julienschmidt
go install

此时在pkg\windows_amd64\github.com\julienschmidt就会有新的文件产生
文件：httprouter.a
这样就能够引入安装的包的api了。
golang中的import name，实际是到GOPATH中去寻找name.a, 使用时是该name.a的源码中生命的package 名字；
```



# 编译问题

```
1.系统编译时 go install abc_name时，系统会到GOPATH的src目录中寻找abc_name目录，然后编译其下的go文件；

 2.同一个目录中所有的go文件的package声明必须相同，所以main方法要单独放一个文件，否则在eclipse和liteide中都会报错；
```

