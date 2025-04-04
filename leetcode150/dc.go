package main

import "fmt"

/*
Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced binary search tree.

Example 1:

Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:

Example 2:

Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.

Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in a strictly increasing order.
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0], Right: nil, Left: nil}
	}
	var recursive func(nums []int, left, right int) *TreeNode

	recursive = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}

		mid := left + (right-left)/2

		root := &TreeNode{
			Val:   nums[mid],                     // 0
			Left:  recursive(nums, left, mid-1),  // r(0, 1)
			Right: recursive(nums, mid+1, right), // r(3, 5)
		}
		return root
	}

	return recursive(nums, 0, len(nums)-1)
}

/*
Given the head of a linked list, return the list after sorting it in ascending order.

Example 1:
Input: head = [4,2,1,3]
Output: [1,2,3,4]
Example 2:

Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]
Example 3:

Input: head = []
Output: []

Constraints:

The number of nodes in the list is in the range [0, 5 * 104].
-105 <= Node.val <= 105

Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.e. constant space)?
*/
func sortList(head *ListNode) *ListNode {

	var merge func(l1, l2 *ListNode) *ListNode

	merge = func(l1, l2 *ListNode) *ListNode {
		tmp := &ListNode{Val: 0, Next: nil}
		moving := tmp
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				moving.Next = l1
				l1 = l1.Next
			} else {
				moving.Next = l2
				l2 = l2.Next
			}
			moving = moving.Next
		}

		if l1 == nil {
			moving.Next = l2
		}

		if l2 == nil {
			moving.Next = l1
		}

		return tmp.Next

	}

	var sort func(head *ListNode) *ListNode

	sort = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}

		current := head
		mid := head
		var tmp *ListNode
		for current != nil && current.Next != nil {
			tmp = mid
			mid = mid.Next
			current = current.Next.Next
		}
		tmp.Next = nil

		l1 := sort(head)
		l2 := sort(mid)
		return merge(l1, l2)

	}

	return sort(head)

}

/*
Given a n * n matrix grid of 0's and 1's only. We want to represent grid with a Quad-Tree.

Return the root of the Quad-Tree representing grid.

A Quad-Tree is a tree data structure in which each internal node has exactly four children. Besides, each node has two attributes:

val: True if the node represents a grid of 1's or False if the node represents a grid of 0's. Notice that you can assign the val to True or False when isLeaf is False, and both are accepted in the answer.
isLeaf: True if the node is a leaf node on the tree or False if the node has four children.

	class Node {
	    public boolean val;
	    public boolean isLeaf;
	    public Node topLeft;
	    public Node topRight;
	    public Node bottomLeft;
	    public Node bottomRight;
	}

We can construct a Quad-Tree from a two-dimensional area using the following steps:

If the current grid has the same value (i.e all 1's or all 0's) set isLeaf True and set val to the value of the grid and set the four children to Null and stop.
If the current grid has different values, set isLeaf to False and set val to any value and divide the current grid into four sub-grids as shown in the photo.
Recurse for each of the children with the proper sub-grid.

If you want to know more about the Quad-Tree, you can refer to the wiki.

Quad-Tree format:

You don't need to read this section for solving the problem. This is only if you want to understand the output format here. The output represents the serialized format of a Quad-Tree using level order traversal, where null signifies a path terminator where no node exists below.

It is very similar to the serialization of the binary tree. The only difference is that the node is represented as a list [isLeaf, val].

If the value of isLeaf or val is True we represent it as 1 in the list [isLeaf, val] and if the value of isLeaf or val is False we represent it as 0.

Example 1:

Input: grid = [[0,1],[1,0]]
Output: [[0,1],[1,0],[1,1],[1,1],[1,0]]
Explanation: The explanation of this example is shown below:
Notice that 0 represents False and 1 represents True in the photo representing the Quad-Tree.

Example 2:

Input: grid = [[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]
Output: [[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
Explanation: All values in the grid are not the same. We divide the grid into four sub-grids.
The topLeft, bottomLeft and bottomRight each has the same value.
The topRight have different values so we divide it into 4 sub-grids where each has the same value.
Explanation is shown in the photo below:

Constraints:

n == grid.length == grid[i].length
n == 2x where 0 <= x <= 6
*/
func construct(grid [][]int) *QuadNode {

	var dc func(x, y, length int) *QuadNode

	dc = func(x, y, length int) *QuadNode {
		
		if length == 1 {
			return &QuadNode{
				Val:         grid[x][y] != 0,
				IsLeaf:      true,
				TopLeft:     nil,
				TopRight:    nil,
				BottomLeft:  nil,
				BottomRight: nil,
			}
		}
		
		root:= &QuadNode{}
		topLeft:= dc(x,y, length/2)
		topRight:= dc(x, y + length/2 , length/2)
		bottomLeft:= dc(x+ length/2, y, length/2)
		bottomRight:= dc(x+length/2, y+length/2, length/2)

		if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf && topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
			root.IsLeaf = true
			root.Val = topLeft.Val
		}else {
			root.IsLeaf = false
			root.Val = false
			root.TopLeft = topLeft
			root.TopRight = topRight
			root.BottomLeft = bottomLeft
			root.BottomRight = bottomRight
		}
		fmt.Println(root.Val)

		return root
	}
	return dc(0, 0, len(grid))
}


/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

 

Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []
 

Constraints:

k == lists.length
0 <= k <= 104
0 <= lists[i].length <= 500
-104 <= lists[i][j] <= 104
lists[i] is sorted in ascending order.
The sum of lists[i].length will not exceed 104.
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) ==0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	var merge func(start, end int) *ListNode

	merge = func(start, end int) *ListNode {
		if start == end {
			return lists[start]
		}
		
		mid:= (start+end)/2 
		listA:= merge(start, mid)
		listB := merge(mid+1, end)
		if listA == nil{
			return listB
		}
		if listB == nil {
			return listA
		}
		/*
		1 -> 4 -> 5
		tmp  listA
		2 -> 3 -> 6
		listB
		*/
		var head *ListNode
		if listA.Val > listB.Val {
			head = listB
			listB = listB.Next
		}else {
			head = listA
			listA = listA.Next
		}
		tmp := head
		for listA != nil && listB!= nil {
			if listA.Val > listB.Val {
				tmp.Next = listB
				listB = listB.Next
			}else {
				tmp.Next = listA
				listA = listA.Next
			}
			tmp = tmp.Next
		} 
		if listA != nil {
			tmp.Next = listA
		}else {
			tmp.Next = listB
		}

		return head		

	}
	return merge(0, len(lists)-1)

}