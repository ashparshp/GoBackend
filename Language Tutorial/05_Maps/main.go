package main

import (
	"fmt"
	"maps"
)

/* Maps are Go's built-in associative data type (sometimes called hashes or dicts in other languages). */
func main() {
	/* Create a map */
	m := make(map[string]int)

	/* Set key/value pairs */
	m["k1"] = 7
	m["k2"] = 13

	/* Print all key/value pairs */
	fmt.Println("map:", m)

	/* Get the value for a key */
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	/* Get the length of a map */
	fmt.Println("len:", len(m))

	/* Delete a key/value pair */
	delete(m, "k2")
	fmt.Println("map:", m)

	/* Clear all key/value pairs */
	clear(m)
	fmt.Println("map:", m)

	/* If the key is not in the map, then the value is the zero value for the value type */
	prs := m["k2"]
	fmt.Println("prs:", prs)

	/* You can also use the optional second return value to determine if the key is present in the map */
	_, prs1 := m["k2"]
	fmt.Println("prs:", prs1)

	/* You can also declare and initialize a new map in the same line */
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	/* You can also iterate over a map */
	for key, value := range n {
		fmt.Println("key:", key, "value:", value)
	}

	/* Check if a key exists in a map */
	k, ok := n["Phones"]
	if ok {
		fmt.Println("key:", k)
	} else {
		fmt.Println("key not found")
	}

	/* Check if two maps are equal */
	m1 := map[string]int{"foo": 1, "bar": 2}
	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("m1 == m2:", maps.Equal(m1, m2))
}