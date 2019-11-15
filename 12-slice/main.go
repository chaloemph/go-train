package main

import "fmt"

func main() {
	number := make([]int, 5)
	fmt.Println(number)
	numberSlice := []int{1, 2, 3}
	fmt.Println(numberSlice)
	numberSlice2 := append(numberSlice, 4, 5)
	fmt.Println(numberSlice2)
	numberSlice3 := numberSlice2[0:2]
	fmt.Println(numberSlice3)
	numberSlice4 := make([]int, 2)
	copy(numberSlice4, numberSlice)
	fmt.Println(numberSlice4)
}
