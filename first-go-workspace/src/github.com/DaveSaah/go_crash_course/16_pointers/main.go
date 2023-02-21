package main

import "fmt"

func main() {
    var a int = 42
    var b *int = &a
    a = 23
    fmt.Println(a, *b)
}
