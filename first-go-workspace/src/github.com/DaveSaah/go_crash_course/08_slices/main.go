package main

import "fmt"

func main() {
    // var nums = []int{1, 2, 3}
    // nums[3] = 5
    // fmt.Println("Capacity", cap(nums))

    // creating slices using make()
    // make(slice_obj, length, capacity)
    // var grades = make([]int, 3, 100)
    // fmt.Printf("Length: %v\nCapacity: %v\n", len(grades), cap(grades))

    someNums := make([]int, 2, 15)
    fmt.Println(someNums)
    fmt.Printf("Length: %v\nCapacity: %v\n", len(someNums), cap(someNums))

    someNums = append(someNums, 2)
    someNums[0] = 34
    someNums = append(someNums, []int{1, 2, 3, 4, 5}...)
    fmt.Println(someNums)
    fmt.Printf("Length: %v\nCapacity: %v\n", len(someNums), cap(someNums))
}
