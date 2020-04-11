/*
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

class Solution
{
public:
    int minimumTotal(vector<vector<int>> &triangle)
    {
        int ret = 0x3f3f3f3f, sz = triangle.size();
        // base case
        for (int i = 1; i < sz; i++)
        {
            triangle[i][0] += triangle[i - 1][0];
            triangle[i][i] += triangle[i - 1][i - 1];
        }
        // triangle[i][j]+= min(triangle[i-1][j-1], triangle[i-1][j])
        for (int i = 2; i < sz; i++)
        {
            for (int j = 1; j < i; j++)
            {
                triangle[i][j] += min(triangle[i - 1][j - 1], triangle[i - 1][j]);
            }
        }
        // select min value
        for (int i = 0; i < triangle[sz - 1].size(); i++)
        {
            ret = min(ret, triangle[sz - 1][i]);
        }
        return ret;
    }
};