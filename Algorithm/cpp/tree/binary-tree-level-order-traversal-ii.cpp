/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution
{
public:
    vector<vector<int>> levelOrderBottom(TreeNode *root)
    {
        vector<vector<int>> res;
        if (root == nullptr)
            return res;
        queue<TreeNode *> Q;
        TreeNode *curNode;
        Q.push(root);
        while (!Q.empty())
        {
            vector<int> levelRes;
            int size = Q.size();
            for (int i = 0; i < size; i++)
            {
                curNode = Q.front();
                Q.pop();
                levelRes.emplace_back(curNode->val);
                if (curNode->left != nullptr)
                    Q.push(curNode->left);
                if (curNode->right != nullptr)
                    Q.push(curNode->right);
            }
            res.emplace_back(levelRes);
        }
        reverse(res.begin(), res.end());
        return res;
    }
};