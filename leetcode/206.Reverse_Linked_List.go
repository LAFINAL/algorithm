package leetcode

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	next := reverseList(head.Next)
	if next != nil {
		var temp int
		temp = head.Val
		head.Val = next.Val
		next.Val = temp
	}

	return head
}
