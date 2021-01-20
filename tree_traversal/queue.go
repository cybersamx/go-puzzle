package main

type Queue struct {
	First  *QueueNode
	Last   *QueueNode
	Length int
}

type QueueNode struct {
	Next  *QueueNode
	Value *TreeNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(value *TreeNode) {
	last := q.Last
	q.Last = &QueueNode{
		Value: value,
	}

	if q.Length == 0 {
		q.First = q.Last
	} else {
		last.Next = q.Last
	}

	q.Length++
}

func (q *Queue) Pop() *TreeNode {
	if q.Length > 0 {
		treeNode := q.First.Value
		q.First = q.First.Next
		q.Length--
		if q.Length == 0 {
			q.Last = q.First
		}

		return treeNode
	}

	return nil
}
