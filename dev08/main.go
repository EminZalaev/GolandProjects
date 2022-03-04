//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import (
	"fmt"
)

func bitToOne(number int64, index int) int64 {
	newNumber := number | (1 << index)
	return newNumber
}

func bitToZero(number int64, index int) int64 {
	newNumber := number &^ (1 << index)
	return newNumber
}

func main() {
	var testNumber1 int64 = 64
	var testNumber2 int64 = 63
	fmt.Printf("Первое число: %b\n", testNumber1)
	fmt.Printf("Устанавливаем второй бит в 1 : %b \n", bitToOne(testNumber1, 2))
	fmt.Printf("Второе число: %b\n", testNumber2)
	fmt.Printf("Устанавливаем второй бит в 0 : %b \n", bitToZero(testNumber2, 2))
}
