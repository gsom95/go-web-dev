package main

import (
	"html/template"
	"os"
)

type User struct {
	Bio template.HTML
}

func main() {
	t, err := template.ParseFiles("cmd/exp/hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Bio: `<script>alert("Haha, you have been h4x0r3d!");</script>`,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
