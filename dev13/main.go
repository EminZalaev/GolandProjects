//Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 5
	b := 6
	fmt.Println(a, b)
	a, b = b, a //в Go можно менят местами две перенменные не создавая третью
	fmt.Println(a, b)
}
