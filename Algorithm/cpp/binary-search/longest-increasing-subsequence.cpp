/*
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4 
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*/
class Solution
{
public:
    int lengthOfLIS(vector<int> &nums)
    {
        if (nums.size() == 0)
            return 0;
        vector<int> d(nums.size() + 1, 0);
        int len = 1;
        d[len] = nums[0];
        for (int i = 1; i < nums.size(); i++)
        {
            if (nums[i] > d[len])
                d[++len] = nums[i];
            else
            {
                int l = 1, r = len, pos = 0;
                while (l <= r)
                {
                    int mid = l + (r - l) / 2;
                    if (nums[i] > d[mid])
                    {
                        pos = mid;
                        l = mid + 1;
                    }
                    else
                    {
                        r = mid - 1;
                    }
                }
                d[pos + 1] = nums[i];
            }
        }
        return len;
    }
};