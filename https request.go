package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Https request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
