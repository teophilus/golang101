package main

import "fmt"

func main() {
	utopianTree(0)
	utopianTree(1)
	utopianTree(2)
	utopianTree(3)
	utopianTree(4)
}

func utopianTree(cycles int) int {
	i := 0
	size := 1
	for i < cycles {
		// growth doubles
		size = size * 2
		i++
		// growth by 1 meter
		size++
		i++
	}

	// account for odd numbers which are half cycles
	if cycles % 2 == 1 {
		size = size - 1
	}

	fmt.Println(cycles, size)
	return size
}