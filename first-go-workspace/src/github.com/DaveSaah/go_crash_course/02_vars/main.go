package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num int = 45
    fmt.Printf("%v, %T\n", num, num)

    var num1 float32
    num1 = float32(num)
    fmt.Printf("%v, %T\n", num1, num1)

    var num2 float32 = 42.5
    fmt.Printf("%v, %T\n", num2, num2)


    var num3 int
    num3 = int(num2)
    fmt.Printf("%v, %T\n", num3, num3)

    var str string
    str = strconv.Itoa(num)	// Itoa converts from an integer into an ascii string
    fmt.Printf("%v, %T\n", str, str)
}
