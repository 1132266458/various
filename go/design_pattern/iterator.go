package main

import "fmt"

type Aggregate interface {
	// 生成迭代器
	Iterator() Iterator
}

//迭代接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BookShelf struct {
	books []*Book
}

// 实现接口方法
func (bs *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{
		bookShelf: bs,
	}
}

func (bs *BookShelf) AddBook(book *Book) {
	bs.books = append(bs.books, book)
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	idx       int
}

func (bs *BookShelfIterator) SetBookShelf(bookShelf *BookShelf) {
	bs.bookShelf = bookShelf
}
//实现接口方法
func (bs *BookShelfIterator) HasNext() bool {
	return bs.idx < len(bs.bookShelf.books)
}
//实现接口方法
func (bs *BookShelfIterator) Next() interface{} {
	book := bs.bookShelf.books[bs.idx]
	bs.idx++
	return book
}

type Book struct {
	Name string
}

func (b *Book) GetName() string {
	return b.Name
}

func main() {
	book1 := &Book{
		Name: "Chinese",
	}
	book2 := &Book{
		Name: "English",
	}
	book3 := &Book{
		Name: "Math",
	}

	// 实例化书架
	bookShelf := &BookShelf{}
	// 添加书籍
	bookShelf.AddBook(book1)
	bookShelf.AddBook(book2)
	bookShelf.AddBook(book3)

	// 获取迭代器
	iter := bookShelf.Iterator()
	// 迭代打印元素
	for iter.HasNext() {
		book := iter.Next().(*Book)
		fmt.Println(book.Name)
	}
}