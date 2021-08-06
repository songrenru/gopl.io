/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		res := head.Next
	} else {
		res := head
	}
	for {
		if head == nil || head.Next == nil {
			break
		}
		if head.Next != nil {
			head.Next.Next, head.Next := head, head.Next.Next
		}
		head = head.Next
	}
	return res
}