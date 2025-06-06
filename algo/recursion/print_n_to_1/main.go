package main

import "fmt"

func main() {
	printTill1(10)
	fmt.Println()

}

func printTill1(n int) {
	fmt.Printf("%d ", n)
	if n > 1 {
		printTill1(n - 1)
	}
}
