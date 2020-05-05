/*
三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。

示例1:

 输入：n = 3 
 输出：4
 说明: 有四种走法
示例2:

 输入：n = 5
 输出：13
提示:

n范围在[1, 1000000]之间

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/three-steps-problem-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

class Solution
{
public:
    int waysToStep(int n)
    {
        long long dp[3] = {1, 1, 2}, mod = 1e9 + 7;
        for (int i = 3; i <= n; i++)
        {
            dp[i % 3] = (dp[i % 3] + dp[(i - 1) % 3] + dp[(i - 2) % 3]) % mod;
        }
        return dp[n % 3];
    }
};