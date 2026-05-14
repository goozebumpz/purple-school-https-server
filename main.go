package main

import "fmt"

type User struct {
	Name string
}

type Named interface {
	setName(name string)
}

func (u *User) setName(name string) {
	u.Name = name
}

func main() {
	var t any = "string"

	if _, ok := t.(string); ok {
		fmt.Println("is string")
	}

	t = User{Name: "Dick"}

	if _, ok := t.(User); ok {
		fmt.Println("is User")
	}

	if named, ok := t.(Named); ok {
		named.setName("Eric")
	}
}
