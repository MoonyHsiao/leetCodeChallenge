package main

import "fmt"

func main() {
	data := 19
	res := isHappy(data)
	fmt.Println(res)
}
func isHappy(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n == 4 || n == 16 || n == 37 || n == 58 || n == 89 || n == 145 || n == 42 || n == 20 {
		return false
	}

	result := sqrt(n)

	return isHappy(result)
}

func sqrt(n int) int {

	sum := 0
	for n >= 10 {
		t := n % 10
		sum += t * t
		n = n / 10
	}
	sum += n * n

	return sum

}
