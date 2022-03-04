package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	array := []int{1, 2, 3, 4, 5, 6}
	go func() {
		for i := 0; i < len(array); i++ {
			ch1 <- array[i]
		}
		close(ch1)
	}()
	go func() {
		for i := 0; i < len(array); i++ {
			ch2 <- array[i] * 2
		}
		close(ch2)
	}()
	for x := range ch2 {
		fmt.Println(x)
	}
}
