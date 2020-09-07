package main

import (
	"fmt"
	"runtime"
	"sync"
)

var err error //for storing errors

var vg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())

	vg.Add(1)
	go foo()
	bar()

	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())
	vg.Wait()
	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())
}

// func (r receiver) identifier(parameters) (return(s)) { code }
func foo() {
	for i := 0; i <= 2; i++ {
		fmt.Println("foo", i)
	}
	vg.Done()
}

func bar() {
	for i := 0; i <= 2; i++ {
		fmt.Println("bar", i)
	}
}

//func (p user) String() string {
//	return fmt.Sprintf("%s %v: %d", p.First, p.Last, p.Age)
//}
