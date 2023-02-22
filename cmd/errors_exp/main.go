package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNotFound = errors.New("not found")
var ErrOtherError = errors.New("something happened")

func A() error {
	// return ErrNotFound
	return ErrOtherError
}

func B() error {
	err := A()
	if err != nil {
		return fmt.Errorf("b: %w", err)
	}
	return nil
}

func main() {
	err := B()
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			log.Fatalln("Not found something:", err)
		}
		log.Fatalln("some error occured:", err)
	}
}
