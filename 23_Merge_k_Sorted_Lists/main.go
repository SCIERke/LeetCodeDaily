package main

import "fmt"

func countListsNull(lists []*ListNode)int {
	cnt := 0;
	for _ , value := range lists {
		if value != nil {
    	cnt++
		}
	}
	return cnt;
}

func mergeKLists(lists []*ListNode) *ListNode {
	tmpList := &ListNode{};
	curr := tmpList;

	if (len(lists) == 0) {
		return curr.Next;
	}

	if (len(lists) == 1) {
		curr.Next = lists[0];
		return curr.Next;
	}

	for (countListsNull(lists) > 1) {
		min_of_lists := int(1e9);
		idx_min_of_lists := -1;

		for i := 0; i < len(lists);i++ {
			list := lists[i];

			if (list != nil) {
				if (list.Val < min_of_lists) {
					min_of_lists = list.Val;
					idx_min_of_lists = i;
				}
			}
		}

		curr.Next = lists[idx_min_of_lists];
		curr = curr.Next;

		lists[idx_min_of_lists] = lists[idx_min_of_lists].Next;
	}

	for i := 0; i < len(lists);i++ {
		if lists[i] != nil {
			curr.Next = lists[i];
		}
	}

	return tmpList.Next;
}

func main() {
	lists := []*ListNode{
		buildList([]int{1, 4 ,5}),
		buildList([]int{1,3,4}),
		buildList([]int{2,6}),
	}

	merged_list := mergeKLists(lists);

	for merged_list != nil {
		fmt.Println(merged_list.Val);

		merged_list = merged_list.Next;
	}
}