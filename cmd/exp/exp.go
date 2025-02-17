package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	t, err := template.ParseFiles("cmd/exp/hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{Name: "John Doe", Age: 42}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
