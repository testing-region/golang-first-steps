package main

import "fmt"

type myStruct struct {
    some int
}

func main() {
    var ms *myStruct = &myStruct{some: 42}
    fmt.Println(ms)

    var ms2 *myStruct = new(myStruct)
    fmt.Println(ms2)
    ms2.some = 42
    fmt.Println(ms2.some)
}
