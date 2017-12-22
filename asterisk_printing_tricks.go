package main

import (
	"fmt"
)

func printAstDescending(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func main() {
	printAstDescending(8)
}