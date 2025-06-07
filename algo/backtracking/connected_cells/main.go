package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main1() {

	result := findLargestRegion([][]int{
		{1, 1, 0, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 1, 1},
	}, 5, 5) // Should return 7
	fmt.Println(result) // Output: 6
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the number of rows (M):")
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Enter the number of columns (N):")
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	matrix := make([][]int, M)
	fmt.Println("Enter the matrix (row by row, space-separated):")
	for i := 0; i < M; i++ {
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			scanner.Scan()
			num, _ := strconv.Atoi(scanner.Text())
			matrix[i][j] = num
		}
	}
	result := findLargestRegion(matrix, M, N)
	fmt.Println("Largest region size:", result) // Output the size of the largest region
}

func fingConnectedCells(matrix [][]int, M, N, r, c int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxSize := 0
	if r < 0 || r >= M || c < 0 || c >= N || matrix[r][c] == 0 {
		return 0
	}
	if matrix[r][c] == 1 {
		// Mark the cell as visited
		matrix[r][c] = 0
		maxSize = 1 // Count this cell
		// Explore all 8 directions
		maxSize += fingConnectedCells(matrix, M, N, r-1, c-1) // Top-left
		maxSize += fingConnectedCells(matrix, M, N, r-1, c)   // Top
		maxSize += fingConnectedCells(matrix, M, N, r-1, c+1) // Top-right
		maxSize += fingConnectedCells(matrix, M, N, r, c-1)   // Left

		maxSize += fingConnectedCells(matrix, M, N, r, c+1)   // Right
		maxSize += fingConnectedCells(matrix, M, N, r+1, c-1) // Bottom-left
		maxSize += fingConnectedCells(matrix, M, N, r+1, c)   // Bottom
		maxSize += fingConnectedCells(matrix, M, N, r+1, c+1) // Bottom-right
	}
	// fmt.Println("fingConnectedCells called with r:", r, "c:", c, "maxSize:", maxSize)

	return maxSize
}

func findLargestRegion(matrix [][]int, M, N int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxSize := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			size := fingConnectedCells(matrix, M, N, i, j)
			if size > maxSize {
				maxSize = size
			}
		}
	}
	return maxSize
}
