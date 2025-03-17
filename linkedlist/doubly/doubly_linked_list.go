package linkedlist

type node struct {
    data int
    next *node
    prev *node
}

type DoublyLinkedList struct {
    head *node
    tail *node
    size uint
}

func NewDoublyLinkedList() *DoublyLinkedList {
    return &DoublyLinkedList{
        head: nil,
        tail: nil,
        size: 0,
    }
}

func (l *DoublyLinkedList) IsEmpty() bool {
    return l.head == nil
}

func (l *DoublyLinkedList) Back() int {
    return l.tail.data
}

func (l *DoublyLinkedList) Front() int {
    return l.head.data
}

func (l *DoublyLinkedList) InsertAtFront(v int) {

    n := &node {
        data: v,
        next: l.head,
        prev: nil,
    }

    if l.IsEmpty() {
        l.tail = n
    } else {
        l.head.prev = n
    }

    l.head = n
    l.size++
}

func (l *DoublyLinkedList) InsertAtBack(v int) {

    n := &node {
        data: v,
        next: nil,
    }

    if l.IsEmpty() {
        l.head = n
    } else {
        l.tail.next = n
        l.tail.prev = l.tail
    }

    l.tail = n
    l.size++
}

func (l *DoublyLinkedList) Find(v int) (*node, bool) {

    curr := l.head

    for curr != nil {
        if curr.data == v {
            return curr, true
        }
        curr = curr.next
    }

    return nil, false
}


