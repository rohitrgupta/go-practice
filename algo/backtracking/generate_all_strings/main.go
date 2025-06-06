package main

import "fmt"

func main() {
	// options := []string{"A", "B", "C"} // Characters to use for generating strings
	options := []string{"0", "1"}
	n := 3 // Length of strings to generate
	result := make([]string, n)
	generateAllStrings(0, options, result)
}

func generateAllStrings(n int, options []string, result []string) {
	for i := 0; i < len(options); i++ {
		result[n] = options[i]
		if n == len(result)-1 {
			printString(result) // Print the generated string
		} else {
			generateAllStrings(n+1, options, result)
		}
	}
}

func printString(result []string) {
	for _, char := range result {
		fmt.Print(char)
	}
	fmt.Println() // Print a newline after the string
}
