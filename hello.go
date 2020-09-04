package main

import (
	"encoding/json"
	"fmt"
)

var err error //for storing errors
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

var people []person

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)

	fmt.Printf("%T\t%T\n", s, bs)
	fmt.Println(s)
	err = json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Printf("Printing a type of error handler: %T\n", err)
		fmt.Println(err)
	}
	fmt.Printf("%+v", people[1])
}

// func (r receiver) identifier(parameters) (return(s)) { code }
