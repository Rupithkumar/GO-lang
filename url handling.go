package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hvvyhv654652105"

func main() {
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	qparams := result.Query()
	fmt.Println(qparams)
	PartsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=Rhaul",
	}
	anotherurl := PartsOfUrl.String()
	fmt.Println(anotherurl)
}
