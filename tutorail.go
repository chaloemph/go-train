package main

import (
	"fmt"
)

func main() {

	str1 := "AAAACTGCTACCGGT"
	str2 := "CTGAATCTACTGCTATTGCAA"
	max := 0
	words := str1

	for i := 0; i < len(str1); i++ {
		for j := i + 1; j < len(str1); j++ {
			for x := 0; x < len(str2); x++ {
				for y := x + 1; y < len(str2); y++ {
					if str1[i:j] == str2[x:y] {
						if len(str1[i:j]) > max {
							max = len(str1[i:j])
							words = str1[i:j]
						}
					}
				}
			}
		}
	}
	fmt.Println(words)
}
