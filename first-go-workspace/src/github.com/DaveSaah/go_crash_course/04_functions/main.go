package main


import "fmt"


func greeting(name string) string {
    return "Hello " + name
}

func getSum(num1, num2, num3 int) int {
    return num1 + num2 + num3
}


func main() {
    fmt.Println(greeting("David"))
    fmt.Println(getSum(3, 9, 6))
}
