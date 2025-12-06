package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    carry := 0

    for l1 != nil || l2 != nil || carry > 0 {
        sum := carry

        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        carry = sum / 10
        sum = sum % 10

        curr.Next = &ListNode{Val: sum}
        curr = curr.Next
    }

    return dummy.Next
}

func main() {
    list1 := buildList([]int{9, 9, 9, 9, 9, 9, 9})
    list2 := buildList([]int{9, 9, 9, 9})

    result_list := addTwoNumbers(list1, list2)

    for result_list != nil {
        fmt.Printf("%d ", result_list.Val)
        result_list = result_list.Next
    }
}
