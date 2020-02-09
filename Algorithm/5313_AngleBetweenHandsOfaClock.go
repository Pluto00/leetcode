package main

/*
给你两个数 hour 和 minutes 。请你返回在时钟上，由给定时间的时针和分针组成的较小角的角度（60 单位制）。

提示：

1 <= hour <= 12
0 <= minutes <= 59
与标准答案误差在 10^-5 以内的结果都被视为正确结果。
给你两个数 hour 和 minutes 。请你返回在时钟上，由给定时间的时针和分针组成的较小角的角度（60 单位制）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/angle-between-hands-of-a-clock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func angleClock(hour int, minutes int) float64 {
	if hour == 12 {
		hour = 0
	}
	var (
		angleHour   float64
		angleMinute float64
		angle       float64
	)
	angleHour = 30*float64(hour) + float64(minutes)/2
	angleMinute = float64(minutes) * 6
	angle = angleMinute - angleHour
	if angle < 0 {
		angle = -angle
	}
	if angle > 180 {
		angle = 360 - angle
	}
	return angle
}
