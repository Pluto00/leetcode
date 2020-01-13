package main

/*
两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。

给出两个整数 x 和 y，计算它们之间的汉明距离。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hamming-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func hammingDistance(x int, y int) int {
	var ans, z = 0, x ^ y
	for z != 0 {
		if z&1 == 1 {
			ans++
		}
		z >>= 1
	}
	return ans
}
