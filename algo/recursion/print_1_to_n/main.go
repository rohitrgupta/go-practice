package main

import "fmt"

func main() {
	printTill1(10)
	fmt.Println()
}

func printTill1(n int) {
	if n > 1 {
		printTill1(n - 1)
	}
	fmt.Printf("%d ", n)
}
