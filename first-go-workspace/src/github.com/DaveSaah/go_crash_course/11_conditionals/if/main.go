package main

import "fmt"

func main() {
    if true {
	fmt.Println("This is true")
    }

    populations := map[string]int {
	"Ghana": 23,
	"Uganda": 33,
	"Senegal": 55,
    }

    if pop, ok := populations["Ghana"]; ok {
	fmt.Println("The population of Ghana is", pop)
    }
}
