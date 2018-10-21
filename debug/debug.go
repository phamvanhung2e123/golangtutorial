package main

import "fmt"

func foo() int {
	n := 1
	return n
}

func bar() int {
	n := 2
	return n
}

func main() {
	n := foo()
	n = bar()
	fmt.Printf("n=%d\n", n)
}