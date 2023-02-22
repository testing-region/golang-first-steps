package main

import "fmt"

func main() {
	greeting := "Hello"
	name := "Stacey"
	sayMsg(greeting, name)
	storeValues(1, 2, 3, 4, 5)
	fmt.Println(returnValues(1, 2, 3, 4, 5))

	d, err := divide(5.0, 0.0)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)
}

func sayMsg(greeting, name string) {
	fmt.Println(greeting, name)
}

func storeValues(values ...int) {
	fmt.Println(values)
}

func returnValues(values ...int) []int {
	return values
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}
