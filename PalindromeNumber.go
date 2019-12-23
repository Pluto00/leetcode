package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return func(x int) int {
		var ans int
		for x > 0 {
			ans = ans*10 + x%10
			x /= 10
		}
		return ans
	}(x) == x
}

func main(){
	fmt.Println(isPalindrome(131))
}