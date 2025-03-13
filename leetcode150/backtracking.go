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
	rs := make([][]int, 0)
	recursion = func(n int, val int, arr []int) {
		if len(arr) == k {
			rs = append(rs, append([]int{}, arr...))
			return
		}
		for i := val; i <= n; i++ {
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
	rs := make([][]int, 0)

	var recursion func(arr []int)
	check := make([]bool, len(nums))

	recursion = func(arr []int) {
		if len(arr) == len(nums) {
			rs = append(rs, append([]int{}, arr...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if !check[i] {
				arr = append(arr, nums[i])
				check[i] = true
				recursion(arr)
				arr = arr[:len(arr)-1]
				check[i] = false
			}
		}
	}
	recursion([]int{})
	return rs

}

/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target.
You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
frequency

	of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
Example 3:

Input: candidates = [2], target = 1
Output: []

Constraints:

1 <= candidates.length <= 30
2 <= candidates[i] <= 40
All elements of candidates are distinct.
1 <= target <= 40
*/
func combinationSum(candidates []int, target int) [][]int {
	rs := make([][]int, 0)

	var recursion func(arr []int, index int, sum int)

	recursion = func(arr []int, index int, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			rs = append(rs, append([]int{}, arr...))
			return
		}

		for i := index; i < len(candidates); i++ {
			sum += candidates[i]
			arr = append(arr, candidates[i])
			recursion(arr, i, sum)
			sum -= candidates[i]
			arr = arr[:len(arr)-1]
		}
	}
	recursion([]int{}, 0, 0)

	return rs
}

/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.
The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once.

Example 1:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
Example 2:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true
Example 3:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false

Constraints:

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board and word consists of only lowercase and uppercase English letters.

Follow up: Could you use search pruning to make your solution faster with a larger board?
*/
func exist(board [][]byte, word string) bool {
	cMax := len(board[0])
	rMax := len(board)

	var recursive func(r, c int, length int) bool
	recursive = func(r, c int, length int) bool {
		if length == len(word) {
			return true
		}
		if r < 0 || c < 0 || r >= rMax || c >= cMax || board[r][c] != word[length] {
			return false
		}

		tmp := board[r][c]
		board[r][c] = '*'

		found := recursive(r+1, c, length+1) || recursive(r-1, c, length+1) || recursive(r, c+1, length+1) || recursive(r, c-1, length+1)

		board[r][c] = tmp
		return found
	}

	for i := 0; i < rMax; i++ {
		for j := 0; j < cMax; j++ {
			if recursive(i, j, 0) {
				return true
			}
		}
	}
	return false

}

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]

Constraints:

1 <= n <= 8

3

(((

*/

func generateParenthesis(n int) []string {
    rs:=make([]string, 0)
	var recursive func(str string, open, close int)

	recursive = func(str string, open, close int) {
		if open == close && open == n {
			rs = append(rs, str)
			return
		}

		if open < n {
			str+="("
			recursive(str, open + 1, close)
			str = str[0: len(str)-1]
		}

		if close < open {
			str+=")"
			recursive(str, open, close + 1)
		}

	}

	recursive("", 0, 0)
	return rs
}


