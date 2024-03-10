package main

import "fmt"

func main() {
	var slice = []int{1, 5, 2, 5, 8, 8, 9}
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	fmt.Printf("%v", slice)
}
