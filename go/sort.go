package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

func main_1() {
    // Define a slice of Person structs
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }

    fmt.Println("Before sorting:", people)

    // Use sort.Slice to sort the slice in place
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })

    fmt.Println("After sorting:", people)
}