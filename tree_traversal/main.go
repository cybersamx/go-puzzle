package main

import "fmt"

type TreeNode struct {
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
	Value  string
}

// Returns the following tree.
//
//       A
//     /   \
//    B     C
//   / \   /
//  D   E F

func newSampleTree() *TreeNode {
	d := TreeNode{Left: nil, Right: nil, Value: "D"}
	e := TreeNode{Left: nil, Right: nil, Value: "E"}
	f := TreeNode{Left: nil, Right: nil, Value: "F"}
	b := TreeNode{Left: &d, Right: &e, Value: "B"}
	c := TreeNode{Left: &f, Right: nil, Value: "C"}
	a := TreeNode{Left: &b, Right: &c, Value: "A"}

	return &a
}

func depthFirstDump(n *TreeNode) {
	if n == nil {
		return
	}

	fmt.Println(n.Value)

	depthFirstDump(n.Left)
	depthFirstDump(n.Right)
}

// Notes
// 1. To do breath first, it would be easy to abstract
// the Tree into levels. For example, the Tree has
// the following levels:
//
//       A     level = 0
//     /   \
//    B     C  level = 1
//   / \   /
//  D   E F    level = 2
//
// 2. Use a queue to iterate sideway.

func breathFirstDump(n *TreeNode) {
	// Create queue to push tree nodes
	var queue = Queue{}
	queue.Push(n)

	for queue.Length > 0 {
		// Get top of the node
		current := queue.Pop()
		fmt.Println(current.Value)

		// Examine node and add to the queue if there are children
		if current.Left != nil {
			queue.Push(current.Left)
		}
		if current.Right != nil {
			queue.Push(current.Right)
		}
	}
}

func main() {
	tree := newSampleTree()

	fmt.Println("Depth first...")
	depthFirstDump(tree)

	fmt.Println("Breadth first...")
	breathFirstDump(tree)
}
