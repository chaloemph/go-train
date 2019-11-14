package main

import "fmt"

func main() {
	color := "red"
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	case "green":
		fmt.Println("color is green")
	default:
		fmt.Println("not know color")
	}
}
