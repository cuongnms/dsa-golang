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
	node5:=&ListNode{Val: 7, Next: nil}
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
	fmt.Println(pairSum(node1))
}