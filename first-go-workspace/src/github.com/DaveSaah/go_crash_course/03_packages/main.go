package main

// To import more than 1 packages, wrap them in parenthesis
// Separate them by spaces or newlines.
import (
    "fmt"
    "math"
    "github.com/DaveSaah/golang-first-steps/03_packages/strutil"
)


func main() {
    fmt.Println("Floor(2.7):", math.Floor(2.7))
    fmt.Println("Ceil(2.7):", math.Ceil(2.7))
    fmt.Println("Sqrt(16):", math.Sqrt(16))
    fmt.Println("Reverse 'hello':", strutil.Reverse("Hello"))
}
