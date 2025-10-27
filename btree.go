package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	left_left := &Node{4, nil, nil}
	left_right := &Node{5, nil, nil}
	right_left := &Node{6, nil, nil}
	right_right := &Node{7, nil, nil}

	right := &Node{3, right_left, right_right}
	left := &Node{2, left_left, left_right}

	root := &Node{1, left, right}
	return &Tree{root}
}

// Print all the nodes of a binary tree using BFS algorithm.
func (t *Tree) String() string {
	output := ""
	queue := NewQueue[*Node]()
	queue.Enqueue(t.root)

	for !queue.IsEmpty() {
		current_level := []int{}
		level_size := queue.Len()

		for range level_size {
			current_node, _ := queue.Next()
			current_level = append(current_level, current_node.Value)
			if current_node.Left != nil {
				queue.Enqueue(current_node.Left)
			}
			if current_node.Right != nil {
				queue.Enqueue(current_node.Right)
			}
		}
		output += fmt.Sprintf("%v\n", current_level)
	}
	return output
}

// MaxDepth return the maximum depth of the tree using DFS algorithm
func (t *Tree) MaxDepth() int {
	return DFS(t.root)	
}

func DFS(root *Node) int {
	if root == nil {
		return 0
	}
	left := DFS(root.Left)
	right := DFS(root.Right)
	return 1 + max(left, right)
}
