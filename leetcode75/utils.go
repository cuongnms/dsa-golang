package main

import "fmt"


type ListNode struct {
	Val  int
	Next *ListNode
}

func (ll *ListNode) PrintList() {
	current := ll
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	// Tạo gốc của cây
	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}

	index := 1 // Bắt đầu từ phần tử thứ 2 trong mảng
	for len(queue) > 0 && index < len(values) {
		// Lấy phần tử đầu tiên trong hàng đợi
		current := queue[0]
		queue = queue[1:]

		// Gán con trái
		if index < len(values) && values[index] != nil {
			current.Left = &TreeNode{Val: values[index].(int)}
			queue = append(queue, current.Left)
		}
		index++

		// Gán con phải
		if index < len(values) && values[index] != nil {
			current.Right = &TreeNode{Val: values[index].(int)}
			queue = append(queue, current.Right)
		}
		index++
	}

	return root
}


func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	q:= make([]*TreeNode,0)
	q = append(q, root)
	for len(q) > 0 {
		node:=q[0]
		q=q[1:]
		if node == nil {
			fmt.Println("nil")
		}else {
			fmt.Println(node.Val)
		}
		if node != nil {
			if node.Right == nil && node.Left == nil {
				continue
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}else {
				q = append(q, nil)
			}
			if node.Right != nil {
				q= append(q, node.Right)
			}else {
				q = append(q, nil)
			}
		}
	}
}
