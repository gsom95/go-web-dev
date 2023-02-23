package main

import "fmt"

func foo(params ...int) {
	fmt.Println("is nil =", params == nil)
	fmt.Println("len:", len(params))
	fmt.Println("params:", params)
	fmt.Println("----")
}

func main() {
	foo()
	foo([]int{}...)
	foo(1, 2, 3)
	foo([]int{4, 5, 6, 7}...)
}
