package main

/*
给定一个整数，将其转化为7进制，并以字符串形式输出。

注意: 输入范围是 [-1e7, 1e7] 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/base-7
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func convertToBase7(num int) string {
	if num < 0 {
		return "-" + convertToBase7(-num)
	}
	var ans []uint8
	for num != 0 {
		ans = append(ans, uint8('0'+num%7))
		num /= 7
	}
	reverseUint8Slice := func(data []uint8) []uint8 {
		for i := 0; i < len(data)>>1; i++ {
			data[i], data[len(data)-i-1] = data[len(data)-i-1], data[i]
		}
		return data
	}
	return string(reverseUint8Slice(ans))
}
