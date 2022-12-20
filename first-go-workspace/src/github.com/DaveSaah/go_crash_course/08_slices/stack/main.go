package main

import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5}

    // append() allows us to add elements onto a stack
    a = append(a, 6)
    fmt.Println(a)

    a = a[:len(a)-1]
    fmt.Println(a)
}
