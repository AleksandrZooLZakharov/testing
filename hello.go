package main

//go get ...things... not imported yet)
import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

var err error //for storing errors
var uns []person

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	//fmt.Println(s)

	err = json.Unmarshal([]byte(s), &uns)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", uns)
	fmt.Println("-------")
	for _, v := range uns {
		fmt.Println(v.First, v.Last, v.Age, "loves to say:")
		for _, v2 := range v.Sayings {
			fmt.Println("\t", v2)
		}
	}
	fmt.Println("\n-----------\t---------\n")
	err := json.NewEncoder(os.Stdout).Encode(uns)
	if err != nil {
		fmt.Println(err)
	}

}

// func (r receiver) identifier(parameters) (return(s)) { code }
