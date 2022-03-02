// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func newSet(a []string) []string {
	createdSet := make(map[string]bool)

	for _, i := range a {
		createdSet[i] = true
	}

	result := make([]string, 0, len(createdSet))
	for i := range createdSet {
		result = append(result, i)
	}

	return result
}

func main() {
	stringSubsequence := []string{"cat", "cat", "dog", "cat", "tree"}

	out := newSet(stringSubsequence)
	fmt.Println(out)
}
