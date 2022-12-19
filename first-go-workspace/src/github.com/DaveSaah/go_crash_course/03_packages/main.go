package main

// To import more than 1 packages, wrap them in parenthesis
// Separate them by spaces
import (
    "fmt"
    "math"
    "github.com/DaveSaah/golang-first-steps/03_packages/strutil"
)


func main() {
   fmt.Println(math.Floor(2.7))
   fmt.Println(math.Ceil(2.7))
   fmt.Println(math.Sqrt(16))
   fmt.Println(strutil.Reverse("Hello"))
}
