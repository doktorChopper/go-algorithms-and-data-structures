package threadsafe

import "sync"

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
    defer s.mu.Unlock()

    return s.top == nil 
}

func (s *ThreadSafeStack) Top() int {
    s.mu.RLock()
    defer s.mu.Unlock()

    return s.top.data
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
