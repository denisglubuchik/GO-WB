package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) Hi() string {
	return "Hi " + h.name
}

type Action struct {
	Human
}

func main() {
	Petya := Human{name: "Petya", age: 20}
	fmt.Println(Petya.Hi())

	action := Action{Human: Petya}
	fmt.Println(action.Hi())
}
