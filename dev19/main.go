//Разработать программу, которая переворачивает
//подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
package main

import "fmt"

func reversedString(strings string) string {
	stringRune := []rune(strings)
	reverseRune := make([]rune, len(stringRune))
	for i := 0; i < len(stringRune); i++ {
		reverseRune[i] = stringRune[len(stringRune)-1-i]
	}
	return string(reverseRune)
}

func main() {
	firstString := "s''123;asd;aweqd"
	secondString := "ввыфвфы1231323э123э"
	fmt.Println(reversedString(firstString))
	fmt.Println(reversedString(secondString))
}
