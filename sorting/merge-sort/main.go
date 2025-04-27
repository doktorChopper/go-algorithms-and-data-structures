package main

import "fmt"

type ListNode struct {
    val     int
    next    *ListNode
}

func split(head *ListNode) *ListNode {

    slow := head
    fast := head

    for fast != nil && fast.next != nil {
        fast = fast.next.next
        if fast != nil {
            slow = slow.next
        }
    }

    tmp := slow.next
    slow.next = nil
    return tmp
}


func merge(first *ListNode, second *ListNode) *ListNode {

    if first == nil {
        return second
    }

    if second == nil {
        return first
    }

    if first.val < second.val {
        first.next = merge(first.next, second)
        return first
    } else {
        second.next = merge(first, second.next)
        return second
    }
}

func MergeSort(head *ListNode) *ListNode {
    
    if head == nil || head.next == nil {
        return head
    }

    second := split(head)

    head = MergeSort(head)
    second = MergeSort(second)

    return merge(head, second)
}

func main() {

    l := &ListNode{3, nil}

    // a := []int{13, 3, 8, 1, 15, 2, 3, 7}

    a := []int{4, 2, 1}

    for i := len(a) - 1; i >= 0; i-- {
        l = &ListNode{val: a[i], next: l}
    }

    tmp := l
    for tmp != nil {
        fmt.Printf("%d ", tmp.val)
        tmp = tmp.next
    }

    fmt.Println()

    sortl := MergeSort(l)
    tmp = sortl
    for tmp != nil {
        fmt.Printf("%d ", tmp.val)
        tmp = tmp.next
    }


    fmt.Println()

}
