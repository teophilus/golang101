package main

import "fmt"
import "math"

func main() {
	fmt.Println(utopianTree2(1))
	fmt.Println(utopianTree2(2))
	fmt.Println(utopianTree2(3))
	fmt.Println(utopianTree2(4))
}

func utopianTree2(cycles int) int {
	
	a := float64((cycles / 2) + 1)
	b := float64(1)

	if cycles%2 == 1 {
		a = float64(((cycles + 1) / 2) + 1)
		b = float64(2)
	}
	
	size := int(math.Pow(2, a) - b)
	
	return size
}