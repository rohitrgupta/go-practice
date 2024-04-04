package integers

func Add(a, b int) int {
	return a + b
}

func EvenOdd(a int) string {
	if a%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func Multiple3And5(a int) int {
	sum := 0
	for i := 1; i < a; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}