package main

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr_size := head;
	curr := head;
	prev_curr := head;

	size := 0;
	for curr_size != nil {
		size++;
		curr_size = curr_size.Next;
	}

	if n == size {
		head = head.Next
		return head;
	}

	for i := 0; i < size - n - 1;i++ {
		prev_curr = prev_curr.Next;
	}

	for i := 0; i < size - n;i++ {
		curr = curr.Next;
	}

	prev_curr.Next = curr.Next;
	curr.Next = nil;

	return head;
}

func main() {
	list1 := buildList([]int{1,2});

	position := 1;
	rs := removeNthFromEnd(list1, position);

	for rs != nil {
		fmt.Println(rs.Val);
		rs = rs.Next;
	}
}
