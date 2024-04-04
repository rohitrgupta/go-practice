package main

func Sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}
	return result
}

func ArrayDiff(input []int, other []int) []int {
	var result []int
	for _, number := range input {
		found := false
		for _, otherNumber := range other {
			if number == otherNumber {
				found = true
				break
			}
		}
		if !found {
			result = append(result, number)
		}
	}
	return result
}