package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	// 获取链路长度 n
	cur := head
	n := 1
	for cur.Next != nil {
		cur = cur.Next
		n++
	}
	if n == 1 && k%n == 0 {
		return head
	}
	cur.Next = head // 组成环链路
	add := n - k%n
	for add > 0 {
		cur = cur.Next
		add--
	}
	res := cur.Next // 头节点变为第n-k个节点
	cur.Next = nil
	return res
}
