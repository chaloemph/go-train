package main

import "fmt"

func main() {
	var number [5]int
	num := [5]float64{1, 2, 3, 4, 5}

	number[3] = 20
	fmt.Println(number)
	fmt.Println(num)

	for _, n := range num {
		fmt.Println(n)
	}

}
