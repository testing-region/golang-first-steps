package main

import "fmt"

func main() {
    s := [3]int{1, 2, 3}
    fmt.Println(s)

    for _, val := range s {
	fmt.Println(val)
    }
}
