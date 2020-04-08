/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
class Solution
{
public:
    string longestPalindrome(string s)
    {
        int len = s.length();
        if (len <= 1)
            return s;
        int start = 0, end = 0;
        int mlen = 0;
        for (int i = 0; i < len; i++)
        {
            mlen = max(expenAroundCenter(s, i, i), mlen);
            mlen = max(expenAroundCenter(s, i, i + 1), mlen);
            if (mlen > end - start + 1)
            {
                start = i - (mlen - 1) / 2;
                end = i + mlen / 2;
            }
        }
        return s.substr(start, mlen);
    }

private:
    int expenAroundCenter(string s, int left, int right)
    {
        while (left >= 0 && right < s.length() && s[left] == s[right])
        {
            left--;
            right++;
        }
        cout << left << " " << right << endl;
        return right - left - 1;
    }
};
