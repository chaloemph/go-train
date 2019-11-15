package main

import "fmt"

// User is structure
type User struct {
	id       int
	name     string
	lastname string
}

func main() {
	var user User
	user.id = 1
	user.name = "amang"
	user.lastname = "nima"
	fmt.Println(user)
	fmt.Println("user 1 : ", user.name)

	user2 := User{
		id:       2,
		name:     "sofia",
		lastname: "nima",
	}
	fmt.Println(user2)
	fmt.Println("user 2 : ", user2.name)
}
