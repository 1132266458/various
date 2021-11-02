package main

import (
	"log"
)

//Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

var manager *PrototypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func main() {
	t1 := manager.Get("t1")

	t2 := t1.Clone()

	if t1 == t2 {
		log.Fatal("error! get clone not working")
	}

	//c := manager.Get("t1").Clone()
	//
	//t1 := c.(*Type1)
	//if t1.name != "type1" {
	//	log.Fatal("error")
	//}
}



func init() {
	manager = NewPrototypeManager()

	t1 := &Type1{
		name: "type1",
	}
	manager.Set("t1", t1)
}