package main

import "fmt"

func Example() {
	dataCh := make(chan string)

	data := <- dataCh

	fmt.Printf("result: %#v\n", data)
}