//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
package main

import "fmt"

type human struct {
	name string
	age  years
	sex  string
}

type action struct {
	human
	lenght int
}

type years int

func (r human) agePow() years {
	return r.age * r.age
}

func (r action) printAge() years {
	return r.human.agePow()
}

func main() {
	user := action{
		human:  human{"Emin", 22, "YES"},
		lenght: 15,
	}
	fmt.Println(user)
	fmt.Println(action.printAge(user))
}
