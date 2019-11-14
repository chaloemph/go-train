package main

import (
	"fmt"
	"time"
)

func main() {
	words := "hello function"
	somthings(words)
	starttime := time.Now()
	sum := plus(1000000)
	fmt.Println(sum)
	now := time.Now()
	fmt.Println(now.Sub(starttime) )
}

func somthings(str string) {
	fmt.Println(str)
}

func plus(n int) int {
	sum := (n / 2 * (2 + (n - 1)))
	return sum
}
