package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type QuadNode struct {
	Val bool
	IsLeaf bool 
	TopLeft *QuadNode
	TopRight *QuadNode
	BottomLeft *QuadNode
	BottomRight *QuadNode
}

func CreateLL(values []interface{}) *ListNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root:= &ListNode{Val: values[0].(int), Next: nil}
	current:= root
	for i:=1; i < len(values); i++ {
		current.Next = &ListNode{values[i].(int), nil}
		current = current.Next
	}

	return root
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
