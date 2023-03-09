package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`usage: learn_bcrypt COMMAND ARGS...
COMMAND: 
	- hash "some text"
	- compare "secret string" 'hash' # use backticks for hash because of hash format`)

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

func hash(password string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error hashing:", err)
		return
	}
	hash := string(hashedBytes)
	log.Println(hash)
}

func compare(password, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("Error:", err)
		return
	}

	log.Println("true")
}
