package main

import "fmt"

/*
Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].
You may return the answer in any order.

Example 1:
Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.

Example 2:
Input: n = 1, k = 1
Output: [[1]]
Explanation: There is 1 choose 1 = 1 total combination.

Constraints:
1 <= n <= 20
1 <= k <= n

1 2 3 4 5 6 7 8
5 6 7 8
6 7 8

val = 6
k = 4
len(arr) = 0
n = 8 
8 - 6 + 1 + len(arr) < k 

*/
func combine(n int, k int) [][]int {
    
	var recursion func(n int, val int, arr []int)  
	rs:= make([][]int,0)
	recursion = func(n int, val int, arr []int) {
		if len(arr) ==k {
			rs = append(rs, append([]int{}, arr...))
            return
		}
		for i:=val; i <=n; i++ {
			arr = append(arr, i)
			recursion(n, i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	recursion(n, 1, []int{})
	fmt.Println(rs)
	return rs
}

/*
Given an array nums of distinct integers, return all the possible 
permutations
. You can return the answer in any order.

Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]
 

Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/
func permute(nums []int) [][]int {
    rs:= make([][]int, 0)

	var recursion func(arr []int) 
	check:= make([]bool, len(nums))

	recursion = func(arr []int) {
		if len(arr) == len(nums) {
			rs = append(rs, append([]int{}, arr...))
			return
		}

		for i:=0 ; i < len(nums);i++ {
			if !check[i] {
				arr = append(arr, nums[i])
				check[i] = true
				recursion(arr)
				arr = arr[:len(arr)-1]
				check[i] =false
			}
		}
	}
	recursion([]int{})
	return rs

}