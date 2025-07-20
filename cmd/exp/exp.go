package main

import (
	"html/template"
	"os"
)

type User struct {
	Name, Bio string
	Times     int
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
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
