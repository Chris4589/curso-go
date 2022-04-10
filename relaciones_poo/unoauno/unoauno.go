package main

import "fmt"

//las estructuras se puede relacionar

type User struct {
	nombre string
	email  string
	status bool
}

type Student struct {
	user User //la relacion variable de tipo usuario diferente de la herencia
	code string
}

func main() {
	user := User{
		nombre: "chris",
		email:  "ch@gmail.com",
		status: true,
	}

	user2 := User{
		nombre: "da",
		email:  "da@gmail.com",
		status: false,
	}

	student := Student{
		user: user,
		code: "1234",
	}

	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(student)
	fmt.Println(student.user.email)
}
