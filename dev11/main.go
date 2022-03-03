//Реализовать пересечение двух неупорядоченных множеств.
package main

import (
	"fmt"
)

func getIntersection(firstArr []string, secondArr []string) []string {
	set1 := make(map[string]byte)

	for _, k := range firstArr {
		set1[k] += 1
	}

	for _, k := range secondArr {
		set1[k] += 2
	}

	result := make([]string, 0, len(firstArr))

	for k, v := range set1 {
		if v == 3 {
			result = append(result, k)
		}
	}

	return result
}

func main() {

	firstArr := []string{"a", "b", "c", "d", "f"}
	secondArr := []string{"c", "d", "e", "f", "a", "b"}

	fmt.Println("Intersection is: ", getIntersection(firstArr, secondArr))
}
