package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var i, j = 0, 0
	for i = 0; i < len(strs[0]); i++ {
		for j = 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[j-1][i] {
				return strs[0][0 : i]
			}
		}
	}
	return strs[0][0 : i]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"dos"}))
}
