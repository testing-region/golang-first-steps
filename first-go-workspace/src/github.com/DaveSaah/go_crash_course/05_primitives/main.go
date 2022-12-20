package main

import "fmt"

func main() {
    var name string = "DaveSaah"
    fmt.Println("I am", name)
    fmt.Printf("%v, %T\n", name[2], name[2])
    fmt.Printf("%v, %T\n", string(name[2]), name[2])

    // string concatenation
    fmt.Println(name + name)

    // strings can be represented as a collection of bytes
    bytes := []byte(name)
    fmt.Println(bytes)
}
