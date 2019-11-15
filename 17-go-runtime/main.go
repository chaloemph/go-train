package main

import "fmt"

func main() {
	var input string
	go fmt.Scanln(&input)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(input)
}
