package main

import (
	"fmt"
)

func hasPairWithSum(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
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
	arr := []int{1, 2, 4, 7, 11, 15}
	target := 15

	if hasPairWithSum(arr, target) {
		fmt.Println("Found a pair!")
	} else {
		fmt.Println("No pair found.")
	}
}
