# 环境变量

**什么是环境变量**

```shell
我们从命令行想要运行一个程序的时候，待运行的程序往往不是在当前目录。

PATH变量就是用于保存可以搜索的目录路径，如果待运行的程序不在当前目录，操作系统便可以去依次搜索PATH变量变量中记录的目录，如果在这些目录中找到待运行的程序，操作系统便可以运行。

以Go开发为例，但我们使用go install命令编译、安装go程序后，可执行文件是会被保存在$GOPATH/bin路径下；那么我们可以把这个路径加入到PATH变量中，这样我们便可以在任意路径中运行go安装的程序。

例：
go install main.go 会在$GOPATH/bin路径下生成main.exe文件， 而main.exe是不存在当前目录下的，如果在在单曲目录运行main.exe文件，就需要将$GOPATH/bin加入到环境变量path中。
```



**linux添加环境变量**

```shell
打开环境变量文件
vi /etc/profile

添加环境变量
export GOROOT=/usr/local/go        ##Golang安装目录
export PATH=$GOROOT/bin:$PATH
export GOPATH=/home/go  ##Golang项目目录

刷新环境变量
source /etc/profile
```



# 查看CUP个数、核数、逻辑个数

```shell
# 总核数 = 物理CPU个数 X 每颗物理CPU的核数 
# 总逻辑CPU数 = 物理CPU个数 X 每颗物理CPU的核数 X 超线程数

# 查看物理CPU个数
cat /proc/cpuinfo| grep "physical id"| sort| uniq| wc -l

# 查看每个物理CPU中core的个数(即核数)
cat /proc/cpuinfo| grep "cpu cores"| uniq

# 查看逻辑CPU的个数
cat /proc/cpuinfo| grep "processor"| wc -l

查看CPU信息（型号）
cat /proc/cpuinfo | grep name | cut -f2 -d: | uniq -c
```



# 查看内存信息

```shell
cat /proc/meminfo
```



# 创建sanba容器

```shell
docker run -it --name samba -p 445:445 -v /home:/mount -d dperson/samba -u "chj;123456" -s "home;/mount/;yes;no;no;all;user"
```



# docker容器开机自启

```shell
docker update --restart=always 容器id
```

