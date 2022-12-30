package main

import "fmt"

type animal struct {
    name string
    origin string
}

type bird struct {
    animal
    canFly bool
    speedKPH float32
}


func main() {
    parrot := bird{}
    parrot.name = "Parrot"
    parrot.origin = "Sky"
    parrot.canFly = true
    parrot.speedKPH = 100

    fmt.Println(parrot.name)

    // using literal syntax
    vulture := bird{
	animal: animal{name: "Vulture", origin: "Trees"},
	canFly: true,
	speedKPH: 23,
    }

    fmt.Println(vulture)
}
