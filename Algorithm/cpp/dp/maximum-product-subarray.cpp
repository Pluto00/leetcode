/*
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字）。

 

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
class Solution
{
public:
    int maxProduct(vector<int> &nums)
    {
        int ans = -0x3f3f3f3f, maxi = 1, mini = 1;
        for (int i = 0; i < nums.size(); i++)
        {
            if (nums[i] < 0)
            {
                maxi ^= mini;
                mini ^= maxi;
                maxi ^= mini;
            }
            maxi = max(maxi * nums[i], nums[i]);
            mini = min(mini * nums[i], nums[i]);
            ans = max(maxi, ans);
        }
        return ans;
    }
};