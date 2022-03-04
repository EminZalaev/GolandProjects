//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
//с использованием конкурентных вычислений.

package main

import "fmt"

func sumCalculate(numbers ...int) int {
	ch := make(chan int)
	defer close(ch)

	for _, number := range numbers {
		go func(n int) {
			ch <- n * n
		}(number)
	}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += <-ch
	}
	return sum
}

func main() {
	data := []int{2, 4, 6, 8, 10}
	fmt.Println(sumCalculate(data...))
}
