package main

import (
	"fmt"
	"test-go/helloworld"
	"test-go/loop"
	"test-go/object"
)

func main() {
	helloworld.Hello()
	loop.While()
	loop.For()
	person := object.NewWithValue("benz", 21)
	println()
	fmt.Print(person.Age)

}
