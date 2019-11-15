package main

import "fmt"

func main() {
	fmt.Println("Enter your number ?")
	var input float64
	fmt.Scanf("%f", &input)

	fmt.Printf("Your number is %.2f \n", input)

}
