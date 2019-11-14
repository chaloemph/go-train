package main

import "fmt"

func main() {
	fmt.Println("Enter you money : ")
	var money int
	fmt.Scanf("%d", &money)
	if money < 1000 {
		fmt.Println("Your not richman")
	} else if money > 10000 {
		fmt.Println("Your are richman ")
	} else {
		fmt.Println("your are normal man ")
	}

}
