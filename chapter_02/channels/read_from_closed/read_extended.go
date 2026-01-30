package main

import "fmt"

func ExampleExtended() {
	dataCh := make(chan string)

	data, isClosed := <- dataCh

	fmt.Printf("result: %#v isCLosed: %t\n", data, isClosed)
}