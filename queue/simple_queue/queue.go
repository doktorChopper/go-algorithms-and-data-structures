package simplequeue


type node struct {
    data int
    next *node
}

type Queue struct {
    head *node
    tail *node
    size uint
}

func NewQueue() *Queue {
    return &Queue {
        head: nil,
        tail: nil,
        size: 0,
    }
}

func (q *Queue) IsEmpty() bool {
    return q.head == nil
}

func (q *Queue) Back() int {
    return q.tail.data
}

func (q *Queue) Front() int {
    return q.head.data
}

func (q *Queue) Enqueue(v int) {

    n := &node {
        data: v,
        next: nil,
    }

    if !q.IsEmpty() {
        q.tail.next = n
    } else {
        q.head = n
    }

    q.tail = n
    q.size++
} 

func (q *Queue) Dequeue() *node {

    if q.IsEmpty() {
        return nil
    }

    ret := q.head
    q.head = q.head.next
    
    if q.head == nil {
        q.tail = nil
    }

    q.size--
    return ret
}
