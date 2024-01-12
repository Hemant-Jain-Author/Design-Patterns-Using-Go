package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "hello-world-go"
	
	// Split at the first occurrence of "-"
	result := strings.SplitN(input, "-", 2)
	
	// Output the results
	fmt.Printf("First part: %s\n", result[0])
	fmt.Printf("Second part: %s\n", result[1])
}