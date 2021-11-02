package main

import (
	"sync"
)

//Singleton 是单例模式类
type Singleton struct{}

var singleton *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {

	//once.Do 里面的方法只会执行一次
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}

const parCount = 100

func main()  {

	/*ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		log.Fatal("instance is not equal")
	}*/


	/*wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}*/
}