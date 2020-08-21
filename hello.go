package main

import "fmt"

func main() {
	fmt.Println("Hello, world from fmt")
	foo()
	fmt.Println("Something there")
}

func foo() {
	fmt.Println("Hello from fmt from foo() ;)")
}
