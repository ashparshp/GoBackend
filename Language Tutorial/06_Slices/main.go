package main

import (
	"fmt"
	"slices"
)

func main() {

	// Uninitialized slice is nil
	fmt.Println("----- Uninitialized Slice -----")
	var nums []int
	fmt.Println("nums:", nums)
	fmt.Println("Is nil:", nums == nil)
	fmt.Println("Length:", len(nums))

	// Initialized slice is not nil
	fmt.Println("\n----- Initialized Slice -----")
	var nums2 = make([]int, 2, 5) // 2 is the initial length of the slice
	fmt.Println("nums2:", nums2)
	fmt.Println("Is nil:", nums2 == nil)
	fmt.Println("Length:", len(nums2))
	fmt.Println("Capacity:", cap(nums2))

	nums2[0] = 1 // assign value to the first element
	nums2[1] = 2 // assign value to the second element
	fmt.Println("After assigning values:", nums2)

	nums2 = append(nums2, 1, 4, 5, 6)
	fmt.Println("After appending elements:", nums2)
	fmt.Println("New Length:", len(nums2))
	fmt.Println("New Capacity:", cap(nums2)) // capacity doubles when the slice is full

	// Slice literal
	fmt.Println("\n----- Slice Literal -----")
	nums3 := []int{} // empty slice
	nums3 = append(nums3, 1, 2)
	fmt.Println("nums3:", nums3)
	fmt.Println("Length:", len(nums3))
	fmt.Println("Capacity:", cap(nums3))

	// Copy slice
	fmt.Println("\n----- Copy Slice -----")
	nums4 := make([]int, len(nums3))
	copy(nums4, nums3)
	fmt.Println("nums4 (copied from nums3):", nums4)

	// Slice operator
	fmt.Println("\n----- Slice Operator -----")
	var nums5 = []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", nums5)
	fmt.Println("nums5[1:3]:", nums5[1:3]) // [2 3]
	fmt.Println("nums5[:3]:", nums5[:3])   // [1 2 3]
	fmt.Println("nums5[1:]:", nums5[1:])   // [2 3 4 5]
	fmt.Println("nums5[:]:", nums5[:])     // [1 2 3 4 5]

	// Slice functions
	fmt.Println("\n----- Slice Functions -----")
	var nums6 = []int{1, 2, 3, 4, 5}
	var nums7 = []int{6, 7, 8, 9, 10}
	fmt.Println("nums6:", nums6)
	fmt.Println("nums7:", nums7)
	fmt.Println("Are nums6 and nums7 equal?", slices.Equal(nums6, nums7)) // false
	fmt.Println("Are nums6 and nums6 equal?", slices.Equal(nums6, nums6)) // true

	// 2D slice
	fmt.Println("\n----- 2D Slice -----")
	var nums8 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("nums8:", nums8)
	fmt.Println("First row:", nums8[0])
	fmt.Println("Element at [0][1]:", nums8[0][1])

	// Iterating over a slice
	fmt.Println("\n----- Iterating Over 2D Slice -----")
	for i, v := range nums8 {
		fmt.Println("Row", i, ":", v)
	}
}