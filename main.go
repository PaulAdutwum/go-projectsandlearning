package main

import (
	"fmt"
)

// hasPairWithSum checks if there exists a pair of numbers that add up to the target
func hasPairWithSum(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			fmt.Printf("Pair found: %d + %d = %d\n", arr[left], arr[right], target)
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}

func main() {
	fmt.Println(" Welcome to the Two Pointer Sum Finder Tool!")

	var n int
	fmt.Print("Enter the number of elements in the sorted array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Enter the sorted elements:")

	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	var target int
	fmt.Print("Enter the target sum you want to find: ")
	fmt.Scan(&target)

	fmt.Println("\n Searching for a pair...")
	found := hasPairWithSum(arr, target)

	if !found {
		fmt.Println(" No pair found that adds up to the target.")
	}
}
