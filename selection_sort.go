package main

import (
	"fmt"
)

func selectionSort(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}
	return nums
}

func main() {
	unsorted := []int{5, 9, 3, 4, 11, 54, 66, 22, 14, 18}
	sorted := selectionSort(unsorted)
	unsortedTwo := []int{-6, -9, 5, 2, 6, 12, 8, 8, 8, 12, 1, 3}
	sortedTwo := selectionSort(unsortedTwo)
	unsortedThree := []int{5, 4}
	sortedThree := selectionSort(unsortedThree)
	fmt.Println(sorted)
	fmt.Println(sortedTwo)
	fmt.Println(sortedThree)
}
