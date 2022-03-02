package main

import (
	"bufio"
	"fmt"
	"os"
)

//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"time"
//)
//
//func writeToChannel(c chan int, x int) {
//	fmt.Println(x)
//	c <- x
//	close(c)
//	fmt.Println(x)
//}
//
//func main() {
//	c := make(chan int)
//	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
//	s, err := strconv.Atoi(text)
//	if err != nil {
//
//	}
//	for i := 0; i < s; i++ {
//		fmt.Println(1)
//	}
//	go writeToChannel(c, 10)
//	time.Sleep(1 * time.Second)
//}

//// main.go

// Process
//worker.PooledWork(allData)

func main() {
	workersCol, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	arr := make([]int, 10, 10)
	i := 0
	var n int
	// считываем числа пока не будет введен 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		arr[i] = n
		i++
	}
	fmt.Println(workersCol)
	fmt.Println(arr)
}
