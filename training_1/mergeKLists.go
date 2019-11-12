/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
	1->4->5,
	1->3->4,
	2->6
]
输出: 1->1->2->3->4->4->5->6
*/

package training_1

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}

	if len(lists) == 0 {
		return nil
	}

	result := lists[0]

	for i := 1; i < len(lists); i++  {
		result = merge(result, lists[i])
	}

	return result
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		list2.Next = merge(list2.Next, list1)
		return list2
	} else {
		list1.Next = merge(list1.Next, list2)
		return list1
	}
}
