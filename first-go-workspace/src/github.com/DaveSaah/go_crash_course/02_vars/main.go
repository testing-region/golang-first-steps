package main

import "fmt"



func main() {
    // Main types of variables:
    // string
    // bool
    // int int8 int16 int32 int64
    // uint uint8 uint32 uint64 uintptr
    // byte - alias for uint8
    // rune - alias for int32
    // float32 float64
    // complex64 complex128

    // Creating variables

    // using var keyword
    var name string = "DaveSaah"
    // Value type is infered from the value.
    // It is not necessary to explicitly state it.
    // Only make explicitly typed variable definitions 
    // when they are not the same as the default
    var fruit = "Orange"

    var age int = 37
    // Every variable defined in go must be used, else an error is raised.

    fmt.Println(name, age)
    fmt.Println(fruit)

    // search and make notes about fmt pkg -> golang.org/pkg/fmt
    fmt.Printf("%T\n", age) 	// %T is a verb that gets the type of a value

    // using const
    // if you use `var`, you can change the variable in the lifecycle of the program
    // if you use `const`, it is immutable.

    const pi = 3.14

    // observation: Defining a const variable without using it does not throw an error.

    fmt.Println(pi)

    // shorthand method: It can only be used within a function.
    name_q := 23
    fmt.Printf("%T\n", name_q)

    // Multiple assignments is possible
    const food, color = "Rice", "Blue"

}
