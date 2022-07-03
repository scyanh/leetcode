package linkedlist

type linkedListChallenge struct {

}

func NewLinkedListChallenge() *linkedListChallenge {
	return &linkedListChallenge{}
}

type ListNode struct {
	Val int
	Next *ListNode
}

// MiddleNode Given the head of a singly linked list, find the middle node of the linked list.
/*
If there are two middle nodes, return the second middle node.
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.
 */
func (linkedListChallenge) MiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
