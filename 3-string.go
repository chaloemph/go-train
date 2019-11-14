package main

import (
	"fmt"
)

func main() {
	name := "amang nima"
	job := " web developer"
	fmt.Println(name)
	fmt.Println(name[0:5])
	fmt.Println(name[1:])
	fmt.Println(name[:5])

	fmt.Println(name, job)
	fmt.Println(name + job)
}
