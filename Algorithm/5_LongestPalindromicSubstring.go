package main

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// O(n^2)
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	var (
		palindromeLeft   = 0
		palindromeRight  = 0
		palindromeLength = 0
		isPalindrome     func(string, int, int) (int, int, int)
	)
	isPalindrome = func(str string, L int, R int) (int, int, int) {
		for L >= 0 && R < len(str) && str[L] == str[R] {
			L--
			R++
		}
		return R - L - 1, L, R
	}
	for i := 0; i < len(s); i++ {
		length, L, R := isPalindrome(s, i, i)
		if length > palindromeLength {
			palindromeLength = length
			palindromeLeft = L + 1
			palindromeRight = R - 1
		}
		length, L, R = isPalindrome(s, i, i+1)
		if length > palindromeLength {
			palindromeLength = length
			palindromeLeft = L + 1
			palindromeRight = R - 1
		}
	}
	return s[palindromeLeft : palindromeRight+1]
}
