package main

import "fmt"

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]


Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1


Follow up: Could you minimize the total number of operations done?
*/
func moveZeroes(nums []int)  {
    if len(nums) == 1 {
		return
	}
	// 2 1 0 0 1 3 0 0 2 
    //     l
	//         r
    // 2 1 1 0 0 3 0 0 2
    //       l
    //           r
	// 2 1 1 3 0 0 0 0 2
	//         l       
	//                 r
	// 2 1 1 3 2 0 0 0 0

	left:=0
	right:=0
	for right < len(nums) {
		if nums[left] == 0 && nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		} else if nums[left] !=0 {
			left++
			right++
		} else {
			right++
		}
					
	}

	fmt.Println(nums)
}

/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) 
of the characters without disturbing the relative positions of the remaining characters. 
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).

 
Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false

Constraints:

0 <= s.length <= 100
0 <= t.length <= 104
s and t consist only of lowercase English letters.
 

Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109, 
and you want to check one by one to see if t has its subsequence. 
In this scenario, how would you change your code?
*/
// abcdefghi
// ade
// abc
// eghi

func isSubsequence(s string, t string) bool{
	if s== "" {
		return true
	}
	if t=="" {
		return false
	}
	j:=0
	for i:= 0; i < len(s); i++ {
		for j < len(t) {
			if s[i] == t[j] {
				j++
				if i==len(s) - 1 {
					return true
				}
				break
			}else {
				j++
			}
		} 
		
	}
	if j >= len(t) {
		return false
	}else {
		return true
	}
	

}



/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

 

Example 1:


Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1
 

Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/
func maxArea(height []int) int {
	left:= 0
	right:= len(height) - 1
	area := 0
	for left < right {
		if area <= findMin(height[left], height[right])*(right-left) {
			area = findMin(height[left], height[right]) * (right - left)
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func findMin(a, b int)	int {
	if a > b {
        return b
    }
    return a
}


/*
You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array. 

Example 1:

Input: nums = [1,2,3,4], k = 5
1 2 3 4
1 4 4 1


Output: 2
Explanation1: Starting with nums = [1,2,3,4]:
- Remove numbers 1 and 4, then nums = [2,3]
- Remove numbers 2 and 3, then nums = []
There are no more pairs that sum up to 5, hence a total of 2 operations.
Example 2:

Input: nums = [3,1,3,4,3], k = 6
Output: 1
Explanation: Starting with nums = [3,1,3,4,3]:
- Remove the first two 3's, then nums = [1,4,3]
There are no more pairs that sum up to 6, hence a total of 1 operation.


Constraints:

1 <= nums.length <= 105
1 <= nums[i] <= 109
1 <= k <= 109
*/
func maxOperations(nums []int, k int) int {
	count:=0
	i:=0
	for i < len(nums)  {
		j:= i+1
		for j < len(nums) {
			if nums[i] + nums[j] == k {
				count++
			}
		}
		
	}
}