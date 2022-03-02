//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func reversedWords(str string) string {
	words := strings.Fields(str)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	words := "snow dog sun"
	slova := "снег собака солнце"

	fmt.Println(reversedWords(words))
	fmt.Println(reversedWords(slova))
}
