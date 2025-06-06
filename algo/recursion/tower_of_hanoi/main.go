package main

import "fmt"

func main() {
	MoveHaoi(5, "A", "B", "C")
}

func MoveHaoi(n int, from, to, temp string) {
	if n == 1 {
		fmt.Println("Move disk 1 from", from, "to", to)
		return
	}
	MoveHaoi(n-1, from, temp, to)
	fmt.Println("Move disk", n, "from", from, "to", to)
	MoveHaoi(n-1, temp, to, from)
}
