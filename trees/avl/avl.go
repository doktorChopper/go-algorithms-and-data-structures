package avl

type node struct {
    data    int
    left    *node
    right   *node
    height  uint
}

func (n *node) getHeight() uint {
    if n == nil {
        return 0
    }
    return n.height
}

func (n *node) balanceFactor() int {
    return  int(n.left.getHeight()) - int(n.right.getHeight())
}

func (n *node) rotateLeft() *node {

    r := n.right
    n.right = r.left
    r.left = n

    n.fixHeight()
    r.fixHeight()

    return r
}

func (n *node) rotateRight() *node {

    l := n.left
    n.left = l.right
    l.right = n

    n.fixHeight()
    l.fixHeight()

    return l
}

func (n *node) fixHeight() {
    if n.left.getHeight() > n.right.getHeight() {
        n.height = n.left.getHeight() + 1
    } else {
        n.height = n.right.getHeight() + 1
    }
}

type AVLTree struct {
    root *node
}

func NewAVLTree() *AVLTree {
    return &AVLTree {
        root: nil,
    }
}

func balance(n *node) *node {

    n.fixHeight()

    if n.balanceFactor() == -2 {
        if n.right.balanceFactor() > 0 {
            n.right = n.right.rotateRight()
        }
        return n.rotateLeft()
    } else if n.balanceFactor() == 2 {
        if n.left.balanceFactor() < 0 {
            n.left = n.left.rotateLeft()
        }
        return n.rotateRight()
    }

    return n
}

func (a *AVLTree) Add(v int) {
    a.root = addHelper(v, a.root)
}

func addHelper(v int, root *node) *node {

    if root == nil {
        return &node {
            data:      v,
            left:   nil,
            right:  nil,
            height: 1,
        }
    }

    if v < root.data {
        root.left = addHelper(v, root.left)
    } else {
        root.right = addHelper(v, root.right)
    }

    return balance(root)
}

func (a *AVLTree) Delete(v int) {
    a.root = deleteHelper(v, a.root)
}

func deleteHelper(v int, root *node) *node {

    if root == nil {
        return nil
    }

    if v > root.data {
        root.right = deleteHelper(v, root.right)
    } else if v < root.data {
        root.left = deleteHelper(v, root.left)
    } else {
        
        if root.left == nil && root.right == nil {
            root = nil
            return root
        } else if root.left == nil && root.right != nil {
            return root.right
        } else if root.left != nil && root.right == nil {
            return root.left
        } else {

            minNode := root.right
            var parentMinNode *node

            for minNode.left != nil {
                parentMinNode = minNode
                minNode = minNode.left
            }

            if parentMinNode != nil {
                parentMinNode.left = minNode.right
            } else {
                root.right = minNode.right
            }
            
            root.data = minNode.data
        }
    }

    return balance(root)
}

func (a *AVLTree) Search(v int) (*node, bool) {
    return searchHelper(v, a.root)
}

func searchHelper(v int, root *node) (*node, bool) {

    if root == nil {
        return nil, false
    }

    if v == root.data {
        return root, true
    } else if v > root.data {
        return searchHelper(v, root.right)
    }
    return searchHelper(v, root.left)
}
