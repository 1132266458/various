package main

import (
	"fmt"
	"log"
)

//API is interface
type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct{}

//Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

//HelloAPI is another API implement
type helloAPI struct{}

//Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main()  {
	/*api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		log.Fatal("Type1 test fail")
	}*/

	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "Hello, Tom" {
		log.Fatal("Type2 test fail")
	}
}
