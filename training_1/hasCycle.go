/*
给定一个链表，判断链表中是否有环。

使用快慢指针来判断
*/

package training_1

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head.Next, head.Next.Next
	
	for {
		if slow == fast {
			return true
		}

		if slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		} else {
			return false
		}
	}
}
