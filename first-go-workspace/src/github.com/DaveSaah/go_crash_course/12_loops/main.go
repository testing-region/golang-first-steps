package main

import "fmt"

func main() {
    for i := 0; i <= 5; i++ {
	fmt.Println(i)
    }

    for x := 0; x <= 5; x+=2 {
	fmt.Println(x)
    }

    for i, j := 0, 0; i <= 5; i, j = i+1, j+1 {
	fmt.Println(i, j)
    }

    // when the value is already initialised in your program
    i := 0
    fmt.Println("Before loop... i =", i)

    for ; i <= 5; i++ {
	fmt.Println(i)
    }

    fmt.Println("After loop... i =", i)

    // Placing the increamenter in the loop
    for i := 0; i < 5; {
	fmt.Println(i)
	i++
    }

    // while loop implementation
    for i < 10 {
	fmt.Println(i)
	i++
    }
}
