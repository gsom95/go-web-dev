package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`usage: learn_bcrypt COMMAND ARGS...
COMMAND: 
	- hash "some text"
	- compare "secret string" "hash"`)

		return
	}

	switch v := os.Args[1]; v {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid commands: %s\n", v)
	}
}

func hash(hashStr string) {
	fmt.Printf("TODO: implement me: %q\n", hashStr)
}

func compare(secret, hash string) {
	fmt.Printf("TODO: implement me: secret string = %q, hash = %q\n", secret, hash)
}
