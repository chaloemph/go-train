package main

import "fmt"

func main() {
	country := make(map[string]string)
	country["th"] = "thailand"
	country["jp"] = "japan"
	fmt.Println(country["th"])
	fmt.Println(country)

	country2 := map[string]string{
		"th": "thailand",
		"jp": "japan",
	}

	fmt.Println(country2)

}
