package main

import "fmt"


type person struct {
    name string
    age int
    siblings []string
}

func main() {
    person1 := person {
	name: "David Saah",
	age: 20,
	siblings: []string{
	    "Linux",
	    "Golang",
	    "Julia",
	    "Python",
	},
    }

    fmt.Println(person1.name)
    fmt.Println(person1.age)
    fmt.Println(person1.siblings)
}
