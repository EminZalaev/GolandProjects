//Реализовать бинарный поиск встроенными методами языка.
package main

import (
	"errors"
	"fmt"
)

func binarSearch(array []int, item int) (int, error) {
	lowKey := 0
	highKey := len(array) - 1
	for lowKey <= highKey {
		mid := (lowKey + highKey) / 2
		guess := array[mid]
		if guess == item {
			return mid, nil
		} else if guess > item {
			highKey = mid - 1
		} else {
			lowKey = mid + 1
		}
	}

	return 0, errors.New("No number in array")
}

func main() {
	// Массив должен быть отсортирован!
	// Худшее время - логарифмическое O(log n); лучшее - константное O(1)
	arr := []int{0, 5, 7, 11, 15, 18, 22}

	result, err := binarSearch(arr, 5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
