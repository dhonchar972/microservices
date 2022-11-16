package model

import "fmt"

type User struct {
	Name string
	Age  int
}

func CreateNewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func (u User) SayHello() {
	fmt.Println("Hello!")
}
