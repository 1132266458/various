package main

import "fmt"

// 统一接口
type List interface {
	Add(idx int, e interface{}) bool
	AddLast(e interface{}) bool
	Remove(idx int) (interface{}, bool)
	RemoveLast() (interface{}, bool)
	Get(idx int) (interface{}, bool)
}

type ArrayList struct {
	ele  []interface{}
	size int
}

func NewArrayList(capacity int) *ArrayList {
	if capacity < 16 {
		capacity = 16
	}
	return &ArrayList{
		ele:  make([]interface{}, 0, capacity),
		size: 0,
	}
}

func ArrayListOf(ele ...interface{}) *ArrayList {
	return &ArrayList{
		ele:  ele,
		size: len(ele),
	}
}

func (a *ArrayList) Add(idx int, e interface{}) bool {
	if idx > a.size {
		return false
	}
	a.ele = append(a.ele, nil)
	copy(a.ele[idx+1:], a.ele[idx:]) // 后移一位, 把idx位置空出来
	a.ele[idx] = e
	a.size++
	return true
}

func (a *ArrayList) AddLast(e interface{}) bool {
	a.ele = append(a.ele, e)
	a.size++
	return true
}

func (a *ArrayList) Remove(idx int) (interface{}, bool) {
	if idx >= a.size {
		return nil, false
	}
	e := a.ele[idx]
	a.ele = append(a.ele[:idx], a.ele[idx+1:]...) // 跳过idx元素
	a.size--
	return e, true
}

func (a *ArrayList) RemoveLast() (interface{}, bool) {
	e := a.ele[a.size-1]
	a.ele = append(a.ele[:a.size-1])
	a.size--
	return e, true
}

func (a *ArrayList) Get(idx int) (interface{}, bool) {
	if idx >= a.size {
		return nil, false
	}
	e := a.ele[idx]
	return e, true
}

func (a *ArrayList) String() string {
	return fmt.Sprintf("size=%v, %v", a.size, a.ele)
}

func main() {
	var list List
	list = NewArrayList(10)
	list.AddLast(1)
	list.AddLast(3)
	list.AddLast(4)
	list.Add(1, 2)
	fmt.Println(list)              // size=4, [1 2 3 4]
	fmt.Println(list.RemoveLast()) // 4 true
	fmt.Println(list)              // size=3, [1 2 3]
	fmt.Println(list.Remove(0))    // 1 true
	fmt.Println(list)              // size=2, [2 3]
	fmt.Println(list.Get(1))       // 3 true

	list = ArrayListOf(1, 2, 3, 4, 5, 6)
	fmt.Println(list)              // size=6, [1 2 3 4 5 6]
	fmt.Println(list.AddLast(7))   // true
	fmt.Println(list)              // size=7, [1 2 3 4 5 6 7]
	fmt.Println(list.Add(1, 100))  // true
	fmt.Println(list)              // size=8, [1 100 2 3 4 5 6 7]
	fmt.Println(list.Remove(1))    // 100 true
	fmt.Println(list)              // size=7, [1 2 3 4 5 6 7]
	fmt.Println(list.RemoveLast()) // 7 true
	fmt.Println(list)              // size=6, [1 2 3 4 5 6]
	fmt.Println(list.Get(2))       // 3 true
}