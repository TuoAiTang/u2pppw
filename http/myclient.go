package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	url := "http://httpbin.org/ip"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	bytes, err:= httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	response, err := http.DefaultClient.Get(url)
	bytes, err = httputil.DumpResponse(response, true)
	fmt.Println(string(bytes))

	var body []byte
	resp.Body.Read(body)
	fmt.Println("------BODY------")
	fmt.Println(string(body))

}
