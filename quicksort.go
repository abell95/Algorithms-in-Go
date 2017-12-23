//TODO

package main

import (
	"fmt"
	"sort"
)

func quickSort(numList []int) {
	//todo
	//return sorted list
}

func main() {
	arr := []int{8, 99, 13, 5, 73, 22, 6, 44, 5, 11}
	sortedArr := quicksort(arr)
	fmt.Println(sortedArr)
	if sort.IntsAreSorted(sortedArr) == true {
		fmt.Println("It's sorted!")
	}
}
