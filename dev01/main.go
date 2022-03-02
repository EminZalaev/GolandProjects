package main

import "fmt"

type Human struct {
	name string
	age  int
	sex  string
}

type Action struct {
	Human  Human
	lenght int
}

func main() {
	user := Human{"Vasya", 23, "Male"}
	fmt.Println(user)
	action := Action{Human: user, lenght: 20}
	fmt.Println(action)
}
