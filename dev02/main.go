//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых
//из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	PrintSquares(numbers...)
}

func PrintSquares(numbers ...int) {
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for _, number := range numbers {
		go func(n int) {
			defer wg.Done()
			fmt.Println(Square(n))
		}(number)
	}
	wg.Wait()
}

func Square(number int) int {
	return number * number
}
