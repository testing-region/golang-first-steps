package main

import "fmt"

func main() {
	s := func(name string) (msg string) {
		return "Hello " + name
	}

	fmt.Println(s("David"))
}
