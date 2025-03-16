package trees

type node struct {
    data    int
    left    *node
    right   *node
}

type BinarySearchTree struct {
    root *node
}

func NewBinarySearchTree() *BinarySearchTree {
    return &BinarySearchTree {
        root: nil,
    }
}

func (b *BinarySearchTree) IsEmpty() bool {
    return b.root == nil
}

func (b *BinarySearchTree) Add(v int) {

    n := &node {
        data:   v,
        left:   nil,
        right:  nil,
    }

    if b.IsEmpty() {
        b.root = n
        return
    }

    curr := b.root
    var parent *node = nil

    for curr != nil {
        parent = curr
        if v <= curr.data {
            curr = curr.left
        } else {
            curr = curr.right
        }
    }

    if v <= parent.data {
        parent.left = n
    } else {
        parent.right = n
    }

}

func (b *BinarySearchTree) Search(v int) (*node, bool) {

    curr := b.root

    for curr != nil {
        if v < curr.data {
            curr = curr.left
        } else if v > curr.data {
            curr = curr.right
        } else {
            return curr, true
        }
    }

    return nil, false
}

func (b *BinarySearchTree) Max() *node {

    curr := b.root

    for curr.right != nil {
        curr = curr.right
    }

    return curr
}

func (b *BinarySearchTree) Min() *node {

    curr := b.root

    for curr.left != nil {
        curr = curr.left
    }

    return curr
}
