package main

import "fmt"

type myStruct struct {
    some int
}

func main() {
    var a int = 42
    var b *int = &a
    fmt.Println(a, *b)
    fmt.Println(&a, b)

    a = 24
    fmt.Println(a, *b)

    *b = 19
    fmt.Println(a, *b)

    d := [3]int {1, 2, 3}
    e := &d[0]
    f := &d[1]
    fmt.Printf("%v, %p, %p\n", d, e, f)

    var ms *myStruct
    ms = &myStruct{some: 42}
    fmt.Println(*ms)
}
