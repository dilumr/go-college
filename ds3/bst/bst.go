package bst

// Body contains the application-specific information stored in a BST node.
type Body interface {
	LessThan(other interface{}) bool
}

type node struct {
	left  *node
	right *node
	body  Body
}

// BinarySearchTree arranges Body data in (current unbalanced) binary search tree form.
type BinarySearchTree struct {
	root *node
}

// New builds a new BST
func New() *BinarySearchTree {
	return &BinarySearchTree{nil}
}

// VisitInOrder traverses the BST in-order, calling back to visitor with each Body.
func (bst *BinarySearchTree) VisitInOrder(visitor func(body interface{})) {
	recursiveInOrder(bst.root, visitor)
}

func recursiveInOrder(current *node, visitor func(body interface{})) {
	if current == nil {
		return
	}
	recursiveInOrder(current.left, visitor)
	visitor(current.body)
	recursiveInOrder(current.right, visitor)
}

// Find returns a body in the BST that not less and not greater than the search body.
func (bst *BinarySearchTree) Find(body Body) (Body, bool) {
	current := bst.root
	for current != nil {
		if body.LessThan(current.body) {
			current = current.left
		} else if current.body.LessThan(body) {
			current = current.right
		} else {
			return current.body, true
		}
	}
	return nil, false
}

func (bst *BinarySearchTree) FindNearest(body Body) (Body, Body, bool, bool) {
	sentinel := &node{}
	lt := sentinel
	gt := sentinel
	current := bst.root
	for current != nil {
		if body.LessThan(current.body) {
			gt = current
			current = current.left
		} else if current.body.LessThan(body) {
			lt = current
			current = current.right
		} else {
			return current.body, current.body, true, true
		}
	}
	return lt.body, gt.body, (lt != sentinel), (gt != sentinel)
}

// Insert incorporates the body into the BST.
// Adds if the body is less or greater than all bodies already in the BST
// Replaces the existing body if one is equal to the one being inserted.
func (bst *BinarySearchTree) Insert(body Body) {
	bst.root = recursiveInsert(bst.root, body)
}

func recursiveInsert(current *node, body Body) *node {
	if current == nil {
		return &node{nil, nil, body}
	}
	if body.LessThan(current.body) {
		// go left
		current.left = recursiveInsert(current.left, body)
		return current
	}
	if current.body.LessThan(body) {
		// go right
		current.right = recursiveInsert(current.right, body)
		return current
	}
	// replace current body
	current.body = body
	return current
}

// --- APIs for internal diagnostics and invariance checks

// NodeCount is the number of nodes in the BST.
func (bst *BinarySearchTree) nodeCount() int {
	return recursiveNodeCount(bst.root)
}

func recursiveNodeCount(current *node) int {
	if current == nil {
		return 0
	}
	return 1 + recursiveNodeCount(current.left) + recursiveNodeCount(current.right)
}
