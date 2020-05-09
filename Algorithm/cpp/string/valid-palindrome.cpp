/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
class Solution
{
public:
    bool isPalindrome(string s)
    {
        string tmp;
        for (char c : s)
        {
            if (islower(c) || isdigit(c))
                tmp += c;
            else if (isupper(c))
                tmp += (c + 32);
        }
        int i = 0, j = tmp.size() - 1;
        while (i < j && tmp[i] == tmp[j])
        {
            i++;
            j--;
        }
        return i >= j;
    }
};
