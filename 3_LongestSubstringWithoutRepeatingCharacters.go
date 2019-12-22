package main

import "fmt"
/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 */
func lengthOfLongestSubstring(s string) int {
	locationOfChar := map[int32]int{}
	last := -1
	ans := 0
	for i, ch := range s {
		if idx, ok := locationOfChar[ch]; ok && idx >= last {
			last = idx
		}
		locationOfChar[ch] = i
		ans = maxi(ans, i-last)
	}
	return ans
}

func maxi(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring(" "))
}
