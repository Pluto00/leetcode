/*
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

 

示例 1：

输入: n = 3
输出: 6
示例 2：

输入: n = 9
输出: 45
 

限制：

1 <= n <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/qiu-12n-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
class Solution
{
public:
    int sumNums(int n)
    {
        return quickMulti(n, n + 1) >> 1;
    }
    int quickMulti(int A, int B)
    { // return A*B
        int ans = 0;
        // -(B&1)=[0, -1], A&0=0, A&-1=A ==> if(B&1!=0) ans+=A;
        B && (ans = (A & -(B & 1)) + quickMulti(A << 1, B >> 1));
        return ans;
    }
};
