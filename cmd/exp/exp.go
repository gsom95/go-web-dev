package main

import (
	"html/template"
	"os"
)

type User struct {
	Name, Bio  string
	Times      int
	Adventures []string
}

func main() {
	t, err := template.ParseFiles("cmd/exp/hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:  "John Doe",
		Bio:   `<script>alert("Haha, you have been h4x0r3d!");</script>`,
		Times: 100500,
		//Adventures: []string{
		//	"Bank heist",
		//	"Solving food shortage",
		//	"Start a war with kiwi bird",
		//	"Finish the war with kiwi birds",
		//},
		//Adventures: []string{},
		Adventures: nil,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
