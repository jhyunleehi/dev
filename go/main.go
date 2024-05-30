package main

import (
	"fmt"
	"sort"
)

type KValue struct {
	Key   string
	Value int
}

func main() {
	// Initialize the map
	myMap := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
		"Aob":     11125,
		"Cob":     11,
		"Dob":     0,
	}

	// Create a slice to hold the key-value pairs
	var kvList []KValue

	// Iterate over the map and append each key-value pair to the slice
	for k, v := range myMap {
		kvList = append(kvList, KValue{Key: k, Value: v})
	}

	// Print the result
	fmt.Println(kvList)

	sort.Slice(kvList,func(i,j int) bool{
		return kvList[i].Value < kvList[j].Value
	})
	for _,v:= range kvList{
		fmt.Printf("[%d][%s]\n", v.Value, v.Key)
	}

	fmt.Printf("%+v",kvList)
}
