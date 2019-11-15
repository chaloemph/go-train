package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println("increment : ", increment())
	fmt.Println("increment : ", increment())

	add := func(x, y int) int {
		sum := x + y
		return sum
	}

	fmt.Println("sum : ", add(5, 5))

}
