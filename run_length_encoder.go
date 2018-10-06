package main

import (
	"fmt"
	"strconv"
)

// ignores all values not a-z, A-Z
func runLengthEncoder(s string) string {
	var current rune
	counter := 0
	result := ""
	for i, x := range s {
		if i == 0 {
			current = x
		}
		if (64 < x && 91 > x) || (96 < x && 123 > x) {
			if (current != x) {
				result += string(current)
				ctrString := strconv.Itoa(counter)
				result += ctrString
				current = x
				counter = 1
			} else {
				counter++
			}
		}
	}
	// now append the last values
	result += string(current)
	ctrString := strconv.Itoa(counter)
	result += ctrString
	return result
}

func main() {
	fmt.Println(runLengthEncoder("aaaabbbbccc")) // returns a4b4c3
	fmt.Println(runLengthEncoder("AAAAaaaaaaaaAAAAAAaaaa")) // returns A4a8A6a4
	fmt.Println(runLengthEncoder("aaa     ,,,,24%@#$%@#$^&@&@@#@#$%@#$% aaaa bbbb")) //returns a7b4
}
