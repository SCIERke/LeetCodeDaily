package main

import "fmt"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy

    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            curr.Next = list1
            list1 = list1.Next
        } else {
            curr.Next = list2
            list2 = list2.Next
        }
        curr = curr.Next
    }

    // ถ้าตัวใดยังเหลือ ต่อท้ายได้เลย
    if list1 != nil {
        curr.Next = list1
    } else {
        curr.Next = list2
    }

    return dummy.Next
}

func main() {
    list1 := buildList([]int{1, 2, 4})
    list2 := buildList([]int{1, 3, 4})

    mergeList := mergeTwoLists(list1, list2)

    for mergeList != nil {
        fmt.Println(mergeList.Val)
        mergeList = mergeList.Next
    }
}