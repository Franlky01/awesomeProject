package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	fmt.Println("This line would  not get Printed")

}

func test() {
	n := func() {
		fmt.Println("Defer in test")
	}
	defer n()
	msg := "good-bye"
	message(msg)
}

func message(msg string) {
	f := func() {
		fmt.Println("Defer in message func ")
	}
	defer  f()
	if msg == "good-bye" {
		panic(errors.New("something went Wrong"))
	}
}