package main

import "fmt"

func main() {
    var comp complex64 = 1 + 10i
    fmt.Println(comp + comp)
    fmt.Println(comp - comp)
    fmt.Println(comp * comp)
    fmt.Println(comp / comp)

    // To get the complex and imaginary parts separately
    fmt.Println(real(comp), imag(comp))

    // Create a complex number using the complex function
    comp1 := complex(2, 3)
    fmt.Println(comp1)
}
