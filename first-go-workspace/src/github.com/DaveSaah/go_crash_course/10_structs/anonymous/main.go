package main

import "fmt"

func main() {
    doctor := struct{name string}{name: "John Doe"}
    fmt.Println(doctor)

    doctor1 := doctor
    doctor1.name = "Tom Verner"
    fmt.Println(doctor1)

    doctor2 := &doctor
    doctor2.name = "Tom Verner"
    fmt.Println(doctor2)

    fmt.Println(doctor)
}
