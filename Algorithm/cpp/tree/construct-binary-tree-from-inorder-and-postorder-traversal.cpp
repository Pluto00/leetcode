/*
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
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
    TreeNode *buildTree(vector<int> &inorder, vector<int> &postorder, int inl, int inr, int pol, int por)
    {
        if (inl > inr || pol > por)
            return nullptr;
        TreeNode *root = new TreeNode(postorder[por]);
        if (pol == por)
            return root;
        int pos = 0;
        for (int i = inl; i <= inr; i++)
        {
            if (inorder[i] == postorder[por])
            {
                pos = i;
                break;
            }
        }
        root->left = buildTree(inorder, postorder, inl, pos - 1, pol, pol + (pos - 1 - inl));
        root->right = buildTree(inorder, postorder, pos + 1, inr, por + (pos - inr), por - 1);
        return root;
    }

public:
    TreeNode *buildTree(vector<int> &inorder, vector<int> &postorder)
    {
        return buildTree(inorder, postorder, 0, inorder.size() - 1, 0, postorder.size() - 1);
    }
};