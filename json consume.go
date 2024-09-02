package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"course name"`
	Price    int
	Platform string `json:"course platform"`
	Password string `json:"-"`
	Duration string
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json consumption")
	Decodejson()
}
func Decodejson() {
	internetcourse := []courses{
		{"ReactJS", 299, "online", "abc123", "12hrs", []string{"web-dev", "js"}},
		{"MERN", 199, "online", "bcd123", "8hrs", []string{"fullstack", "database"}},
		{"Angular", 399, "online", "cde123", "20hrs", []string{"web-dev", "js"}},
		{"HTML", 99, "online", "cfg123", "12hrs", nil},
	}
	result, err := json.MarshalIndent(internetcourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
	jsondata := []byte(`
	 	{
                "course name": "Angular",
                "Price": 399,
                "course platform": "online",
                "Duration": "20hrs",
                "tags": ["web-dev","js"]
        }
	`)
	checkvalid := json.Valid(jsondata)
	//type1
	var netcourse courses
	if checkvalid {
		fmt.Println("The json data is valid ")
		json.Unmarshal(jsondata, &netcourse)
		fmt.Printf("%#v\n", netcourse)
	} else {
		fmt.Println("The json data is not valid")
	}
	//type 2
	var myonlinecourse map[string]interface{}
	json.Unmarshal(jsondata, &myonlinecourse)
	fmt.Printf("%#v\n", myonlinecourse)
	for v, k := range myonlinecourse {
		fmt.Printf("The key %v value is %v and the type is %T \n", v, k, k)
	}
}
