package main

import "fmt"

func utopianTree(size int, cycles int) int {
	i := 0
		for i < cycles {
			size = size * 2
			i = i + 1
			size = size + 1
			i = i + 1
		}

		fmt.Println(i, size)

	return size
}

func main() {
	utopianTree(1, 0)
	utopianTree(1, 2)
	utopianTree(1, 4)
}