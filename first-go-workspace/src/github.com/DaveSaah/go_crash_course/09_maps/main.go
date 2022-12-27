package main

import "fmt"

func main() {
    statePopulations := map[string] int {
	"state1" : 123,
	"state2" : 345,
	"state3" : 567,
    }

    fmt.Println(statePopulations)
    fmt.Println(statePopulations["state1"])

    statePopulations["Dave"] = 9999
    statePopulations["state1"] = 12233
    fmt.Println(statePopulations)

    // using a key that does not exist
    value, ok := statePopulations["states"]
    fmt.Println(value, ok)
}
