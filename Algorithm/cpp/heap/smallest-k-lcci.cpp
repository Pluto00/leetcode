/*
设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。

示例：

输入： arr = [1,3,5,7,2,4,6,8], k = 4
输出： [1,2,3,4]
提示：

0 <= len(arr) <= 100000
0 <= k <= min(100000, len(arr))

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/smallest-k-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
class Solution
{
private:
    void swap(vector<int> &arr, int i, int j)
    {
        int tmp = arr[i];
        arr[i] = arr[j];
        arr[j] = tmp;
    }

public:
    vector<int> smallestK(vector<int> &arr, int k)
    {
        int n = arr.size();
        if (n <= k)
            return arr;
        vector<int> ans(k);
        buildMinHeap(arr, n);
        for (int i = 0; i < k; i++)
        {
            ans[i] = arr[0];
            swap(arr, 0, --n);
            heapify(arr, 0, n);
        }
        return ans;
    }

    void buildMinHeap(vector<int> &arr, int n)
    {
        for (int i = (n - 2) / 2; i >= 0; i--)
        {
            heapify(arr, i, n);
        }
    }

    void heapify(vector<int> &arr, int i, int n)
    {
        if (i >= n)
            return;
        int lc = i * 2 + 1;
        int rc = i * 2 + 2;
        int smaller = i;
        if (lc < n && arr[lc] < arr[smaller])
            smaller = lc;
        if (rc < n && arr[rc] < arr[smaller])
            smaller = rc;
        if (smaller != i)
        {
            swap(arr, i, smaller);
            heapify(arr, smaller, n);
        }
    }
};