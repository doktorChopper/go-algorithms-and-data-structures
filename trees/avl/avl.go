package avl

type node struct {
    data    int
    left    *node
    right   *node
    height  uint
}

func (n *node) getHeight() uint {
    return n.height
}

func (n *node) balanceFactor() int {
    return int(n.left.getHeight()) - int(n.right.getHeight())
}

func (n *node) rotateLeft() *node {

    b := n.right
    n.right = b.left
    b.left = n

    return n
}

func (n *node) rotateRight() *node {

    b := n.left
    n.left = b.right
    b.right = n

    return b
}

type AVLTree struct {
    root *node
}

func NewAVLTree() *AVLTree {
    return &AVLTree {
        root: nil,
    }
}

