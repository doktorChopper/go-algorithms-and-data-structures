package linkedlist

import "fmt"


type node struct {
    data int
    next *node
}

type SinglyLinkedList struct {
    head *node
    size uint
}

func NewSinglyLinkedList() *SinglyLinkedList {
    return &SinglyLinkedList{
        head: nil,
        size: 0,
    }
}

func (l *SinglyLinkedList) Front() int {
    return l.head.data
}

func (l *SinglyLinkedList) IsEmpty() bool {
    return l.head == nil
}

func (l *SinglyLinkedList) Size() uint {
    return l.size
}

func (l *SinglyLinkedList) InsertAtFront(v int) {

    n := &node {
        data: v,
        next: l.head,
    }

    l.head = n
    l.size++
}

func (l *SinglyLinkedList) InsertAtBack(v int) {

    n := &node {
        data: v,
        next: nil,
    }

    if l.IsEmpty() {
        l.head = n
        return
    }

    curr := l.head

    for curr.next != nil {
        curr = curr.next
    }

    curr.next = n
    l.size++
}

func (l *SinglyLinkedList) DeleteAtFront() {
    if !l.IsEmpty() {
        l.head = l.head.next
    }
    l.size--
}

func (l *SinglyLinkedList) DeleteAtBack() {

    curr := l.head

    for curr.next.next != nil {
        curr = curr.next
    }

    curr.next = nil
    l.size--
}

func (l *SinglyLinkedList) Find(v int) (*node, bool) {
    
    curr := l.head

    for curr != nil {
        if curr.data == v {
            return curr, true
        }
        curr = curr.next
    }

    return nil, false
}

func (l *SinglyLinkedList) Display() {

    curr := l.head

    for curr != nil {
        fmt.Printf("%d ", curr.data)
        curr = curr.next
    }

    fmt.Println()
}

func (l *SinglyLinkedList) Reverse() {

    curr := l.head
    next := curr
    var prev *node = nil

    for curr != nil {

        next = curr.next
        curr.next = prev
        prev = curr
        curr = next
    }

    l.head = prev
}

