package main

import (
	"dasar-golang-belajar-2025/model"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	var username string
	var password string
	day := "siang"
	if day == "siang" {
		username = "fikri"
		password = "abbas"
	} else {
		username = "agus"
		password = "agus"
	}
	// object User to take function CreateUser
	User := model.User{}
	User.CreateUser(username, password)
}
