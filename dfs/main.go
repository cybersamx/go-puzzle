package main

import "fmt"

type Tree struct {
	Left   *Tree
	Right  *Tree
	Parent *Tree
	Value  string
}

// Returns the following tree.
//
//       A
//     /   \
//    B     C
//   / \   /
//  D   E F

func newTree() *Tree {
	d := Tree{Left: nil, Right: nil, Value: "D"}
	e := Tree{Left: nil, Right: nil, Value: "E"}
	f := Tree{Left: nil, Right: nil, Value: "F"}
	b := Tree{Left: &d,  Right: &e,  Value: "B"}
	c := Tree{Left: &f,  Right: nil, Value: "C"}
	a := Tree{Left: &b,  Right: &c,  Value: "A"}

	return &a
}

func depthFirstDump(n *Tree) {
	if n == nil {
		return
	}

	fmt.Println(n.Value)

	depthFirstDump(n.Left)
	depthFirstDump(n.Right)
}

func main() {
	tree := newTree()

	depthFirstDump(tree)
}
