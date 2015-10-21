package main

import "fmt"

func main() {
	utopianTree(1, 4)
}

func utopianTree(size int, cycles int) (int, int) {
	i := 0
	for i < cycles {
		// growth doubles
		size = size * 2
		i = i + 1
		// growth by 1 meter
		size = size + 1
		i = i + 1
	}

	fmt.Println(cycles, size)
	return i, size
	
}