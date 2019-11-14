package main

import "fmt"

func main() {
	somethings(9,8,7,6,5)

}

func somethings(args ...int) {
	for _, i := range args {
		fmt.Println(i)
	}

	for i := range args {
		fmt.Println(args[i])
	}
}
