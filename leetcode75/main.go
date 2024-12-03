package main

import "fmt"


func main() {
	fmt.Println("-------------result--------------")
	// fmt.Println(mergeAlternately("ab", "defg"))

	// fmt.Println(gcdOfStrings("ababab", "abab"))
	// fmt.Println(isCommonSubString("abc", "ab"))
	// fmt.Println(kidsWithCandies([]int{2,3,5,1,3}, 3))
	// fmt.Println(canPlaceFlowers([]int{1,0,0,0,1}, 1))
	// fmt.Println(reverseVowels("Camel"))
	// fmt.Println(reverseWords(" a good  example  "))
	// fmt.Println(productExceptSelf([]int{-1,1,0,3,-3}))
	// fmt.Println(increasingTriplet([]int{5,4,3,2,1}))
	//fmt.Println(compress([]byte{'a','a', 'b', 'b', 'b', 'b', 'c', 'a', 'a'}))
	// moveZeroes([]int{1,2,0,2,4,0,0,6 })
	// fmt.Println(isSubsequence("abc", "ahbgdc"))
	// fmt.Println(maxOperations([]int{0, 4,1,2,5,6,9,3}, 4))
	// fmt.Println(findMaxAverage([]int{9,7,3,5,6,2,0,8,1,9}, 6))
	// fmt.Println(maxVowels("abciiidef", 3))
	// fmt.Println(longestOnes([]int{1,1,1,0,0,0,1,1,1,1,0}, 2))
	// fmt.Println(largestAltitude([]int{-5,1,5,0,-7}))
	// fmt.Println(pivotIndex([]int{1,7,3,6,5,6}))
	// fmt.Println(findDifference([]int{1,2,3}, []int{2,4,6}))
	// fmt.Println(uniqueOccurrences([]int{1,2,2,1,1,3}))
	// fmt.Println(closeStrings("abbzzca", "babzzcz"))
	// fmt.Println(removeStars("leet**cod*e"))
	// fmt.Println(asteroidCollision([]int{-2, -1, 1 , 2}))
	// fmt.Println(decodeString("3[a2[c]]"))

	// obj:=Constructor()
	// fmt.Println(obj.Ping(1))
	// fmt.Println(obj.Ping(100))
	// fmt.Println(obj.Ping(3000))
	// fmt.Println(obj.Ping(3001))
	// fmt.Println(obj.Ping(3002))
	// fmt.Println(obj.Ping(3004))

	// fmt.Println(predictPartyVictory("DRDRR"))
	node1:=&ListNode{Val: 1, Next: nil}
	node2:=&ListNode{Val: 2, Next: nil}
	node3:=&ListNode{Val: 3, Next: nil}
	node4:=&ListNode{Val: 4, Next: nil}
	node5:=&ListNode{Val: 5, Next: nil}
	node6:=&ListNode{Val: 6, Next: nil}
	// node7:=&ListNode{Val: 6, Next: nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	// node4.Next = nil
	node4.Next = node5
	node5.Next = nil
	node6.Next = nil
	// node6.Next = node7
	// deleteMiddle(node1)
	// oddEvenList(node1)
	// node:= reverseList(node1)
	// fmt.Println(pairSum(node1))

	tnode1:=&TreeNode{Val: 1, Left: nil, Right: nil}
	tnode2:=&TreeNode{Val: 2, Left: nil, Right: nil}
	tnode3:=&TreeNode{Val: 3, Left: nil, Right: nil}
	tnode4:=&TreeNode{Val: 4, Left: nil, Right: nil}
	tnode5:=&TreeNode{Val: 5, Left: nil, Right: nil}
	tnode6:=&TreeNode{Val: 6, Left: nil, Right: nil}
	tnode7:=&TreeNode{Val: 7, Left: nil, Right: nil}
	tnode8:= &TreeNode{Val: 8, Left: nil, Right: nil}

	tnode1.Left = tnode2
	tnode1.Right = tnode3
	tnode3.Left = tnode4
	tnode3.Right = tnode5

	tnode4.Left = tnode6
	tnode6.Left = tnode7

	tnode8.Left = tnode3
	tnode8.Right = tnode4
	// fmt.Println(maxDepth(tnode1))
	// fmt.Println(maxDepthIter(tnode1))
	// fmt.Println(leafSimilar(tnode1, tnode8))
	// fmt.Println(pathSum(tnode1, 6))

	//0 1 
	//2 3 4 5
	// addTwoNumbers(node1, node5).Print()
	// swapPairs(node1).Print()
	// fmt.Println(myPow(2.0000, -2))
	// reverseListRecursive(node1).Print()
	// reorderList(node1)
	// removeElements(node1, 4).Print()
	// reverseByHeadRecursive(nil, node1)
	// reverseByTailRecursive(nil, node1).Print()
	// fmt.Println(isPalindrome(node1))
	// fmt.Println(lastRemaining(9))
	// fmt.Println(predictTheWinner([]int{2,4,55,6,8}))

	zig1:=&TreeNode{Val: 1, Left: nil, Right: nil}
	zig2:=&TreeNode{Val: 2, Left: nil, Right: nil}
	zig3:=&TreeNode{Val: 3, Left: nil, Right: nil}
	zig4:=&TreeNode{Val: 4, Left: nil, Right: nil}
	zig5:=&TreeNode{Val: 5, Left: nil, Right: nil}
	zig6:=&TreeNode{Val: 6, Left: nil, Right: nil}
	zig7:=&TreeNode{Val: 7, Left: nil, Right: nil}
	zig8:=&TreeNode{Val: 8, Left: nil, Right: nil}
	// zig9:=&TreeNode{Val: 9, Left: nil, Right: nil}
	// zig10:=&TreeNode{Val: 10, Left: nil, Right: nil}
	zig1.Right = zig2
	zig2.Left = zig3
	zig2.Right = zig4
	zig4.Left = zig5
	zig4.Right = zig6
	zig5.Right = zig7
	zig7.Right = zig8

	values := []interface{}{
		5,3,6,2,4,nil,7,
	}
	root:=CreateTree(values)
	// fmt.Println(longestZigZag(root))
	// fmt.Println(rightSideView(root))
	// fmt.Println(rightSideViewRecursive(root))
	// fmt.Println(maxLevelSum(root))
	// fmt.Println(searchBST(root, 2))
	// printTree(root)
	printTree(deleteNode(root, 3))
	
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
