/*
给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？

以任意顺序返回这两个数字均可。

示例 1:

输入: [1]
输出: [2,3]
示例 2:

输入: [2,3]
输出: [1,4]
提示：

nums.length <= 30000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-two-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
class Solution
{
public:
    vector<int> missingTwo(vector<int> &nums)
    {
        int ans = 0, n = nums.size() + 2;
        for (int i = 1; i <= n; i++)
            ans ^= i;
        for (int num : nums)
            ans ^= num;
        int diff = ans & -ans, one = 0;
        for (int i = 1; i <= n; i++)
        {
            if (i & diff)
            {
                one ^= i;
            }
        }
        for (int num : nums)
        {
            if (num & diff)
            {
                one ^= num;
            }
        }
        return {one, one ^ ans};
    }
};