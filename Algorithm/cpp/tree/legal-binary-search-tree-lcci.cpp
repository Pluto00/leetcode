/*
实现一个函数，检查一棵二叉树是否为二叉搜索树。

示例 1:
输入:
    2
   / \
  1   3
输出: true
示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/legal-binary-search-tree-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
class Solution {
public:
    bool isValidBST(TreeNode* root) 
    {
        if(root==nullptr) return true;
        return isValidBSThelp(root->left, LONG_MIN, root->val) && isValidBSThelp(root->right, root->val, LONG_MAX);
    }
    bool isValidBSThelp(TreeNode* root, long long minVal,long long maxVal)
    {
        if(root==nullptr) return true;
        if(root->val <= minVal || root->val >= maxVal) return false;
        bool lt=isValidBSThelp(root->left, minVal, root->val);
        bool rt=isValidBSThelp(root->right, root->val, maxVal);
        return lt&&rt;
    }
};