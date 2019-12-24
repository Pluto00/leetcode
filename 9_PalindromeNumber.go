package main

import "fmt"
/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

进阶:

你能不将整数转为字符串来解决这个问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
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