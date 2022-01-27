//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых
//из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calculate(param [5]int16) { //функция для расчета квадратов
	for i := 0; i < 5; i++ { //используем цикл для возведения каждого элемента в квадрат
		param[i] = param[i] * param[i] //и сразу же выводим
		fmt.Println(param[i])
		amt := time.Duration(rand.Intn(250)) //ставим рандомную задержку
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	array := [5]int16{2, 4, 6, 8, 10} //добавляем данный нам по условию массив
	go calculate(array)               //выполняем конкурентную функцию
	var input string                  //добавление переменной для stdout
	fmt.Scanln(&input)                //вывод квадратов через stdout
}
