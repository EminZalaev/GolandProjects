// Удалить i-ый элемент из слайса.
package main

import "fmt"

func removeElementFromSlice(arr []int, index int) []int {
	return append(arr[0:index], arr[index+1:]...) //при помощи append
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(removeElementFromSlice(arr, 3))
}
