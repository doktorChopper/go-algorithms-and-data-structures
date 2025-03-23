package splay

type node struct {
    data    int
    left    *node
    right   *node
    parent  *node
}

type SplayTree struct {
    root *node
}

func NewSplayTree() *SplayTree {
    return &SplayTree {
        root: nil,
    }
 }

func (s *SplayTree) IsEmpty() bool {
    return s.root == nil
}

func (s *SplayTree) zig(n *node) {
    p := n.parent
    l := n.left

    if p != nil {
        if p.left == n {
            p.left = l
        } else {
            p.right = l
        }
    }

    n.left = l.right
    if l.right != nil {
        l.right.parent = n
    }

    l.right = n
    n.parent = l
    l.parent = p
}

func (s *SplayTree) zag(n *node) {

    p := n.parent
    r := n.right

    if p != nil {
        if p.left == n {
            p.left = r
        } else {
            p.right = r
        }
    }

    n.right = r.left
    if r.left != nil {
        r.left.parent = n
    }

    r.left = n
    n.parent = r
    r.parent = p
}

func (s *SplayTree) splay(n *node) *node {

    for n.parent != nil {
        if n == n.parent.left {
            if n.parent.parent == nil {
                s.zig(n.parent)
            } else if n.parent == n.parent.parent.left {
                s.zig(n.parent.parent)
                s.zig(n.parent)
            } else {
                s.zig(n.parent)
                s.zag(n.parent)
            }
        } else {
            if n.parent.parent == nil {
                s.zag(n.parent)
            } else if n.parent == n.parent.parent.right {
                s.zag(n.parent.parent)
                s.zag(n.parent)
            } else {
                s.zag(n.parent)
                s.zig(n.parent)
            }
        }
    }

    return n
}

func (s *SplayTree) Delete(v int) bool {

    del, flag := s.Search(v)

    if !flag {
        return false
    }

    if del.left != nil {
        del.left.parent = nil
    }
    s.root = s.splay(findMax(del.left))
    s.root.right = del.right

    if del.right != nil {
        del.right.parent = s.root
    }
    return true
}

func (s *SplayTree) Max() *node {
    return findMax(s.root)
}

func findMax(root *node) *node {

    for root.right != nil {
        root = root.right
    }

    return root
}

func (s *SplayTree) Add(v int) {

    n := &node {
        data:   v,
        left:   nil,
        right:  nil,
        parent: nil,
    }

    if s.IsEmpty() {
        s.root = n
        return
    }

    curr := s.root
    var parentCurr *node

    for curr != nil {
        parentCurr = curr
        if v < curr.data {
            curr = curr.left
        } else {
            curr = curr.right
        }
    }

    if v <= parentCurr.data {
        parentCurr.left = n
    } else {
        parentCurr.right = n
    }

    n.parent = parentCurr

    s.root = s.splay(n)
}

func (s *SplayTree) Search(v int) (*node, bool) {

    curr := s.root
    var prev *node

    for curr != nil {
        prev = curr
        if v < curr.data {
            curr = curr.left
        } else if v > curr.data {
            curr = curr.right
        } else if v == curr.data {
            return s.splay(curr), true
        }
    }

    s.splay(prev)
    return nil, false
}







