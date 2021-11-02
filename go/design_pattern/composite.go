package main

import "fmt"

//Node 节点
type Node interface {
	//添加后续节点
	Add(child Node)
	//space 表示打印名字前，有几个空格
	Print(space int)
}

type File struct {
	name string
}

//Add
func (t *File) Add(child Node) {
	//do nothing
}

//Print
func (t *File) Print(space int) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(t.name)
}
type Folder struct {
	name  string
	child []Node
}

//Add add file or add folder
func (t *Folder) Add(child Node) {
	t.child = append(t.child, child)
}

//Print
func (t *Folder) Print(space int) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(t.name)
	//传入下级目录
	space++
	for _, v := range t.child {
		v.Print(space)
	}
}

func main() {
	//根目录D盘
	diver := Folder{name: "D:"}
	//两个文件
	f1 := File{name: "1.txt"}
	f2 := File{name: "test.txt"}
	//文件夹
	fd1 := Folder{name: "文档"}
	ff1 := File{name: "2.txt"}
	ff2 := File{name: "3.txt"}
	ff3 := File{name: "4.txt"}
	fd1.Add(&ff1)
	fd1.Add(&ff2)
	fd1.Add(&ff3)
	//追加到D文件夹中
	diver.Add(&f1)
	diver.Add(&f2)
	diver.Add(&fd1)
	//打印目录层次
	diver.Print(0)
	return
}