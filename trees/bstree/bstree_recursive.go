package bstree

type BinarySearchTreeRecursion struct {
    root *node
}

func NewBinarySearchTreeRecursion() *BinarySearchTreeRecursion {
    return &BinarySearchTreeRecursion {
        root: nil,
    }
}

func (b *BinarySearchTreeRecursion) IsEmpty() bool {
    return b.root == nil
}

func (b *BinarySearchTreeRecursion) Add(v int) {
    b.root = addHelper(v, b.root)
}

func addHelper(v int, root *node) *node {
    if root == nil {
        return &node {
            data:   v,
            left:   nil,
            right:  nil,
        }
    }

    if v < root.data {
        root.left = addHelper(v, root.left)
    } else {
        root.right = addHelper(v, root.right)
    }

    return root
}

func (b *BinarySearchTreeRecursion) Search(v int) (*node, bool) {
    return searchHelper(v, b.root)
}

func searchHelper(v int, root *node) (*node, bool) {
    if root == nil {
        return nil, false
    }

    if v == root.data {
        return root, true
    } else if v < root.data {
        return searchHelper(v, root.left)
    } 

    return searchHelper(v, root.right)
}

func (b *BinarySearchTreeRecursion) Min() *node {
    return minHelper(b.root)
}

func minHelper(root *node) *node {
    if root.left == nil {
        return root
    }

    return minHelper(root.left)
}

func (b *BinarySearchTreeRecursion) Max() *node {
    return maxHelper(b.root) 
}

func maxHelper(root *node) *node {
    if root.left == nil {
        return root
    }

    return maxHelper(root.right)
}

func (b *BinarySearchTreeRecursion) Delete(v int) {
 
}

func deleteHelper(v int, root *node) *node {
    if root == nil {
        return nil
    }

    if v < root.data {
        root.left = deleteHelper(v, root.left)
    } else if v > root.data {
        root.right = deleteHelper(v, root.right)
    } else {
        if root.left == nil && root.right == nil {
            root = nil
            return nil
        } else if root.left == nil && root.right != nil {
            root = root.right
        } else if root.left != nil && root.right == nil {
            root = root.left
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

    return root
}
