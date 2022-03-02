package main

import "fmt"

func main() {

	char := []byte("aabbccd")
	var j int
	for i := 0; i < len(char)-1; i++ {
		if char[i] == (char[i+1]) {
			j++
		}
		fmt.Println(char[i])
	}
}
