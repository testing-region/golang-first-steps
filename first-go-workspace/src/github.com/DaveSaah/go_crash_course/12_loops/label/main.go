package main

import "fmt"

func main() {
    Loop:
	for i := 1; i <= 3; i++ {
	    for j := 1; j <= 3; j++ {
		val := i * j
		fmt.Println(val)
		if val >= 3 {
		    break Loop
		}
	    }
	}
}
