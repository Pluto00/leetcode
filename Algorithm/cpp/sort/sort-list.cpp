/*
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */

#define dd(x) cout << #x << "=" << x << ", "
#define de(x) cout << #x << "=" << x << endl

class Solution
{
public:
    ListNode *sortList(ListNode *head)
    {
        if (head == nullptr || head->next == nullptr)
            return head;
        ListNode *slow = head;
        ListNode *fast = head->next;
        while (fast != nullptr && fast->next != nullptr)
        {
            slow = slow->next;
            fast = fast->next->next;
        }
        ListNode *tmp = slow->next;
        slow->next = nullptr;
        ListNode *lt = sortList(head);
        ListNode *rt = sortList(tmp);
        ListNode *newHead = new ListNode(0);
        ListNode *ret = newHead;
        while (lt != nullptr && rt != nullptr)
        {
            if (lt->val < rt->val)
            {
                newHead->next = lt;
                lt = lt->next;
            }
            else
            {
                newHead->next = rt;
                rt = rt->next;
            }
            newHead = newHead->next;
        }
        newHead->next = lt == nullptr ? rt : lt;
        return ret->next;
    }
};