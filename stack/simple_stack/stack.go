package simplestack

type node struct {
    data int
    next *node
}

type Stack struct {
    top *node
}

func NewStack() *Stack {
    return &Stack {
        top: nil,
    }
}

func (s *Stack) Top() int {
    return s.top.data
}

func (s *Stack) IsEmpty() bool {
    return s.top == nil
}

func (s *Stack) Push(v int) {
    n := &node {
        data: v,
        next: s.top,
    }

    s.top = n
}

func (s *Stack) Pop() (*node, bool) {
    
    if s.top == nil {
        return nil, false
    }

    ret := s.top
    s.top = s.top.next
    return ret, true
}
