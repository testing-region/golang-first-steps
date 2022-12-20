package main

import "fmt"

func main() {
    //grades := [3]int{98, 99, 97}
    grades_ := [...]int{98, 99, 87, 85, 86}
    var students [5]string

    fmt.Println("Grades:", grades_)
    fmt.Println("Students:", students)


    // assigning values to an existing array
    students[0] = "David"
    students[1] = "Saah"
    students[2] = "Orange"
    fmt.Println("Student #1:", students[0])
    fmt.Println("Student #2-3:", students[1:3])
    fmt.Println(len(students))

    var identityMatrix [3][3]int = [3][3]int {
	[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1},
    }

    var identityMatrix_ [2][2]int 
    identityMatrix_[0] = [2]int{1, 0}
    identityMatrix_[1] = [2]int{0, 1}
    fmt.Println(identityMatrix)
    fmt.Println(identityMatrix_)
}
