package threadsafe

import (
	"fmt"
	"sync"
)

type node struct {
    data    int
    next    *node
}

type ThreadSafeStack struct {
    top     *node
    mu      sync.RWMutex
}

func NewThreadSafeStack() *ThreadSafeStack {
    return &ThreadSafeStack {
        top:    nil,
        mu:     sync.RWMutex{},
    }
}

func (s *ThreadSafeStack) IsEmpty() bool {
    
    s.mu.RLock()
    defer s.mu.RUnlock()

    return s.top == nil 
}

func (s *ThreadSafeStack) Top() (int, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    if s.top == nil {
        return -1, false
    }

    return s.top.data, true
}

func (s *ThreadSafeStack) Push(v int) {

    s.mu.Lock()
    defer s.mu.Unlock()

    n := &node {
        data: v,
        next: s.top,
    }

    s.top = n
}

func (s *ThreadSafeStack) Pop() (*node, bool) {

    s.mu.Lock()
    defer s.mu.Unlock()

    if s.top == nil {
        return nil, false
    }

    ret := s.top
    s.top = s.top.next

    return ret, true
}

func (s *ThreadSafeStack) Display() {
    s.mu.RLock()
    defer s.mu.RUnlock()

    curr := s.top
    for curr != nil {
        fmt.Println(curr.data)
        curr = curr.next
    }
    fmt.Println()
}
