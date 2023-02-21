package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
	Bio  string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	user := User{
		Name: "John Kek",
		Bio:  `<script>alert("Haha, you have been h4x0r3d!");</script>`,
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		log.Fatalln(err)
	}
}
