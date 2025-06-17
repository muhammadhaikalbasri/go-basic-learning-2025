package model

import "fmt"

type User struct{}

func (user *User) CreateUser(username string, password string) {
	fmt.Printf("create user %v with password %v\n", username, password)
	commit()
}

func commit() {
	println("commit")
}
