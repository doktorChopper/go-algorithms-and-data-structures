package arraystack

import "errors"

var ErrStackUnderflow = errors.New("stack underflow")

type Stack struct {
    top uint
    a   []int
}

func NewStack() *Stack {
    return &Stack {
        top:    0,
        a:      make([]int, 0),
    }
}

func (s *Stack) IsEmpty() bool {
    return s.top == 0
}

func (s *Stack) Top() int {
    return s.a[s.top]
}
 
func (s *Stack) Push(v int) {
    s.a = append(s.a, v)
    s.top++
}

func (s *Stack) Pop() (int, error) {

    if s.top == 0 {
        return -1, ErrStackUnderflow 
    }

    s.top--

    return s.a[s.top + 1], nil
}


