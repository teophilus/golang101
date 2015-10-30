package main

import "fmt"

func main() {
	fmt.Println(utopianTree(0))
	fmt.Println(utopianTree(1))
	fmt.Println(utopianTree(2))
	fmt.Println(utopianTree(3))
	fmt.Println(utopianTree(4))
}

func utopianTree(cycles int) int {
	i := 0
	size := 1
	for i < cycles {
		// growth doubles
		size *= 2
		i++
		// growth by 1 meter
		size++
		i++
	}

	// account for odd numbers which are half cycles
	if cycles % 2 == 1 {
		size = size - 1
	}

	return size
}