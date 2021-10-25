# go build go install

**go build**

go build 入口是一个main包，有main包才能生产exe文件，一个mian包里只能有一个唯一的main方法。

go build产生的exe文件是在当前的项目主目录生成。



**go install**

go install 的作用有两步：第一步是编译导入的包文件，所有导入的包文件编译完才会编译主程序；第二步是将编译后生成的可执行文件放到 bin 目录下（$GOPATH/bin），编译后的包文件放到 pkg 目录（$GOPATH/pkg）。

