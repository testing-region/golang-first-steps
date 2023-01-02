package main

import "fmt"

func main() {
    number := 50
    guess := 80

    if guess < number {
	fmt.Println("Too low")
    } else if guess == number {
	fmt.Println("You got it right")
    } else {
	fmt.Println("Too high")
    }

    fmt.Println(!true)
}
