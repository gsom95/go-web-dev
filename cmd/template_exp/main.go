package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	user := User{"John Kek"}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		log.Fatalln(err)
	}
}
