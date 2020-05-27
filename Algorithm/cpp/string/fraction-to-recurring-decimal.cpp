/*
给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。

如果小数部分为循环小数，则将循环的部分括在括号内。

示例 1:

输入: numerator = 1, denominator = 2
输出: "0.5"
示例 2:

输入: numerator = 2, denominator = 1
输出: "2"
示例 3:

输入: numerator = 2, denominator = 3
输出: "0.(6)"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fraction-to-recurring-decimal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
#define dd(x) cout << #x << "=" << x << ", "
#define de(x) cout << #x << "=" << x << endl

class Solution
{
public:
    string fractionToDecimal(int numerator, int denominator)
    {
        if (numerator == 0)
            return "0";
        if (denominator == 0)
            return "";
        string res;
        long long num = abs((long long)numerator);
        long long den = abs((long long)denominator);
        if ((numerator < 0) ^ (denominator < 0))
            res.append("-");
        res.append(to_string(num / den));
        long long remainder = num % den;
        if (remainder != 0)
            res.append(".");
        map<long long, int> re2pos;
        while (remainder != 0)
        {
            remainder *= 10;
            if (re2pos.find(remainder) != re2pos.end())
            {
                res.insert(re2pos[remainder], "(");
                res.append(")");
                break;
            }
            re2pos[remainder] = res.size();
            res.append(to_string(remainder / den));
            remainder %= den;
        }
        return res;
    }
};