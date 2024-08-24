package digital_root

import "fmt"

func Repeat(char string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += char
	}
	return result
}

func VovelCount(str string) int {
	count := 0
	for _, char := range str {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++
		}
	}
	return count
}

func IntToString(num int) string {
	sign := ""
	result := ""
	if num < 0 {
		sign = "-"
		num = -num
	}
	for num > 0 {
		result = fmt.Sprint(num%10) + result
		num = num / 10
	}
	return sign + result
}

func DigitalRoot(num int) int {
	if num < 10 {
		return num
	}
	sum := 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return DigitalRoot(sum)
}
