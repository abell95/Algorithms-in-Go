package main

import (
	"fmt"
)

func printAstAscending(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func printAstDescending(n int) {
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func main() {
	printAstAscending(5)
	printAstDescending(5)
}
