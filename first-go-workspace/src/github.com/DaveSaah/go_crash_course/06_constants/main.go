package main

import "fmt"


const (
    a = iota
    b 
    c 
)

func main() {
    fmt.Printf("%v\n", a)
    fmt.Printf("%v\n", b)
    fmt.Printf("%v\n", c)
}
