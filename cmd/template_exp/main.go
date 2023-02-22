package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
	Info Info
}

type Info struct {
	Bio string
	Age int
}

type Items []int

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	user := User{
		Name: "John Kek",
		Info: Info{
			Bio: `<script>alert("Haha, you have been h4x0r3d!");</script>`,
			Age: 33,
		},
	}
	items := Items{1, 2, 3, 4}
	err = t.Execute(os.Stdout, struct {
		User
		Items
		EmptyItems Items
		MapItems   map[int]any
	}{
		User:       user,
		Items:      items,
		EmptyItems: nil,
		MapItems:   map[int]any{1: 'a', 2: "kekekek"},
	})
	if err != nil {
		log.Fatalln(err)
	}
}
