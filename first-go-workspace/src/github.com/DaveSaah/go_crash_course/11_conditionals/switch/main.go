package main

import "fmt"

func main() {
    switch 2 {
    case 1:
	fmt.Println("one")
    case 2:
	fmt.Println("two")
    default:
	fmt.Println("Nothing")
    }

    switch {
	case 3 == 3:
	    fmt.Println("true")
	case 4 > 5:
	    fmt.Println("false")
    }

    // fallthrough
    switch {
    case 5 < 10:
	fmt.Println("less than 10")
	fallthrough
    case 5 == 20:
	fmt.Println("equal to 20")
    }

    // type switches
    var i interface{} = map[string] int{}

    switch i.(type) {
    case int:
	fmt.Println("This is an integer")
    case string:
	fmt.Println("This is a string")
    case rune:
	fmt.Println("This is a rune")
    default:
	fmt.Printf("Value is %T\n", i)
    }
}
