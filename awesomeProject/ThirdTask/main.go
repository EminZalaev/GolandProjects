//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
//с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calculate(param [5]int16) { //функция для расчета квадратов
	var sum int16
	for i := 0; i < 5; i++ { //используем цикл для возведения каждого элемента в квадрат
		sum += param[i] * param[i]           //и суммируем их
		amt := time.Duration(rand.Intn(250)) //ставим рандомную задержку
		time.Sleep(time.Millisecond * amt)
	}
	fmt.Println(sum)
}

func main() {
	array := [5]int16{2, 4, 6, 8, 10} //добавляем данный нам по условию массив
	go calculate(array)               //выполняем конкурентную функцию
	var input string                  //добавление переменной для stdout
	fmt.Scanln(&input)                //вывод квадратов через stdout
}
