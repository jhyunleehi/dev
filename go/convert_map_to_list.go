package main

import "fmt"

type KeyValue struct {
	Key   string
	Value int
}

func main() {
	// Initialize the map
	myMap := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}

	// Create a slice to hold the key-value pairs
	var kvList []KeyValue

	// Iterate over the map and append each key-value pair to the slice
	for k, v := range myMap {
		kvList = append(kvList, KeyValue{Key: k, Value: v})
	}

	// Print the result
	fmt.Println(kvList)
}
