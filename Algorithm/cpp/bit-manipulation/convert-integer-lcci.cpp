/*
整数转换。编写一个函数，确定需要改变几个位才能将整数A转成整数B。

示例1:

 输入：A = 29 （或者0b11101）, B = 15（或者0b01111）
 输出：2
示例2:

 输入：A = 1，B = 2
 输出：2
提示:

A，B范围在[-2147483648, 2147483647]之间

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-integer-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

class Solution
{
public:
    int convertInteger(int A, int B)
    {
        // return bitset<32> (A^B).count();
        int c = A ^ B;
        int cnt = 0;
        for (int i = 0; i < 32; i++)
        {
            if ((c >> i) & 1)
            {
                cnt += 1;
            }
        }
        return cnt;
    }
};