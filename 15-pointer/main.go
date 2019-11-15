package main

import "fmt"

func main() {
	number := 10
	fmt.Println("number is ", number)
	fmt.Println("address of number is ", &number)

	var p *int
	p = &number
	fmt.Println("p is ", p)
	fmt.Println("address of p is ", &p)

	*p = 50
	fmt.Println("number is ", number)
	fmt.Println("address of number is ", &number)

}
