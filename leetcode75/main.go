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
	node1:=&ListNode{Val: 0, Next: nil}
	node2:=&ListNode{Val: 1, Next: nil}
	node3:=&ListNode{Val: 2, Next: nil}
	node4:=&ListNode{Val: 3, Next: nil}
	node5:=&ListNode{Val: 4, Next: nil}
	node6:=&ListNode{Val: 5, Next: nil}
	// node7:=&ListNode{Val: 6, Next: nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
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
	tnode1.Left = tnode2
	tnode1.Right = tnode3
	tnode3.Left = tnode4
	tnode3.Right = tnode5

	tnode4.Left = tnode6
	tnode6.Left = tnode7

	tnode8:= &TreeNode{Val: 8, Left: nil, Right: nil}
	tnode8.Left = tnode3
	tnode8.Right = tnode4
	// fmt.Println(maxDepth(tnode1))
	// fmt.Println(maxDepthIter(tnode1))
	fmt.Println(leafSimilar(tnode1, tnode8))
}


/*
Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.
Return the number of good nodes in the binary tree.

Example 1:

Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.
Example 2:

Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
Example 3:

Input: root = [1]
Output: 1
Explanation: Root is considered as good.
 
Constraints:

The number of nodes in the binary tree is in the range [1, 10^5].
Each node's value is between [-10^4, 10^4].
*/
func goodNodes(root *TreeNode) int {
    
	rs := countRercusive(root, root.Val)
	return rs

}

func countRercusive(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}
	count :=0
	
	if max <= root.Val {
		max = root.Val
		count++
	}

	return count + countRercusive(root.Left, max) + countRercusive(root.Right, max)

}