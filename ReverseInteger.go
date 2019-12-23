package main

import "fmt"

func reverse(x int) int {
	var ans, flag = 0, 1
	if x < 0 {
		x = -x
		flag = -1
	}
	for x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if ans > 1<<31-1 || ans < - 1>>31 {
		return 0
	}
	return flag * ans
}

func main() {
	fmt.Println(reverse(120))
}
