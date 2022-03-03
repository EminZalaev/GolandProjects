//Разработать программу, которая проверяет, что все символы в строке уникальные
//(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//	aabcd — false

package main

import (
	"fmt"
	"strings"
)

func checkoutString(stringNumber string) bool {
	m := make(map[string]bool)
	stringNumber = strings.ToLower(stringNumber) //т.к регистр не важен всю строку в нижний регистр

	for _, i := range stringNumber {
		if m[string(i)] == true { //мапа для уникальности и проверки
			return false
		}
		m[string(i)] = true
	}

	return true
}

func main() {
	firstString := "abcd"
	secondString := "abCdefAaf"
	thirdString := "aabcd"
	fmt.Println(checkoutString(firstString))
	fmt.Println(checkoutString(secondString))
	fmt.Println(checkoutString(thirdString))
}
