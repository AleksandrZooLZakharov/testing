package main

import (
	"fmt"
	"sort"
)

var err error //for storing errors

type person struct {
	first string
	age   int
}

type ByAge []person

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)

	sort.Sort(ByAge(people))

	fmt.Println(people)
}

// func (r receiver) identifier(parameters) (return(s)) { code }
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].first < a[j].first }
