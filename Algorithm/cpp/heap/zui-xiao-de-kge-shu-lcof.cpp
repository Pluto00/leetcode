/*
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

 

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]
 

限制：

0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
#define dd(x) cout<<#x<<"="<<x<<", "
#define de(x) cout<<#x<<"="<<x<<endl

class Heap {
private:
    vector<int> data;
    void swap(int i, int j) {
        int tmp = data[i];
        data[i] = data[j];
        data[j] = tmp;
    }    
public:

    void insert(int val) {
        data.push_back(val);
        int idx = data.size()-1;
        int fa  = (idx-1)/2;
        while(fa>=0) {
            if(data[idx]<data[fa]) {
                swap(idx, fa);
                idx = fa;
                fa  = (idx-1)/2;
            } else break;
        }
    }

    void pop() {
        data[0] = data[data.size()-1];
        data.pop_back();
        int maxIdx = data.size()-1;
        int smaller;
        for(int fa=0;fa<=maxIdx;fa=smaller) {
            int lc = fa*2+1, rc=fa*2+2;
            if(rc>maxIdx) {
                if(lc==maxIdx&&data[fa]>data[lc]) {
                    swap(fa,lc);
                }
                break;
            }
            smaller = data[lc] < data[rc]? lc:rc;
            if(data[fa]>data[smaller]) swap(fa, smaller);
            else break;
        }
    }

    int top() {
        return data[0];
    }

   
};

class Solution {
public:
    vector<int> getLeastNumbers(vector<int>& arr, int k) {
        Heap minHeap;
        vector<int> res(k);
        for(int val:arr) {
            minHeap.insert(val);
        }
        for(int i=0;i<k;i++){
            res[i] = minHeap.top();
            minHeap.pop();
        }
        return res;
    }
};