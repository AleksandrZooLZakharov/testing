package main

//go get ...things... not imported yet)
import (
	"fmt"
	"sort"
)

type user struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type ByAge []user
type ByLast []user

var err error //for storing errors

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	fmt.Println("RAW---------------------------------")
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, vari := range v.Sayings {
			fmt.Printf("\t%v\n", vari)
		}

	}

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	fmt.Println(users)
	fmt.Println("SORT INNER LEVEL Alphab---------------------------------")
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, vari := range v.Sayings {
			fmt.Printf("\t%v\n", vari)
		}

	}

	sort.Sort(ByLast(users))
	fmt.Println(users)
	fmt.Println("SORT UPPER LEVEL byAge---------------------------------")
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, vari := range v.Sayings {
			fmt.Printf("\t%v\n", vari)
		}

	}
}

// func (r receiver) identifier(parameters) (return(s)) { code }
func (u ByAge) Len() int           { return len(u) }
func (u ByAge) Swap(x, y int)      { u[x], u[y] = u[y], u[x] }
func (u ByAge) Less(x, y int) bool { return u[x].Age < u[y].Age }

func (u ByLast) Len() int           { return len(u) }
func (u ByLast) Swap(x, y int)      { u[x], u[y] = u[y], u[x] }
func (u ByLast) Less(x, y int) bool { return u[x].Last < u[y].Last }

//func (p user) String() string {
//	return fmt.Sprintf("%s %v: %d", p.First, p.Last, p.Age)
//}
