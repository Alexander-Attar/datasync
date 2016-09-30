package main

import "fmt"

func main() {
	X := []string{"1", "3", "4"}
	Y := []string{"1", "2", "3"}

	fmt.Println(compare(X, Y))
	fmt.Println(compare(Y, X))
}

func compare(X, Y []string) []string {
	counts := make(map[string]int)
	var total int
	for _, val := range X {
		counts[val] += 1
		total += 1
	}
	for _, val := range Y {
		if _, ok := counts[val]; ok {
			counts[val] -= 1
			total -= 1
			if counts[val] <= 0 {
				delete(counts, val)
			}
		}
	}

	fmt.Println(counts)
	difference := make([]string, total)
	i := 0
	for val, count := range counts {
		for j := 0; j < count; j++ {
			difference[i] = val
			i++
		}
	}
	return difference
}
