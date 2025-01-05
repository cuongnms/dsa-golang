package main

import (
	"fmt"
)

/*
The Tribonacci sequence Tn is defined as follows:
T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
Given n, return the value of Tn.

Example 1:

Input: n = 4
Output: 4
Explanation:
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4
Example 2:

Input: n = 25
Output: 1389537

Constraints:

0 <= n <= 37
The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31 - 1.
*/
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	var recursive func (n int) int
	memo:= make([]int, n+1)
	memo[0] = 0
	memo[1] = 1
	memo[2] = 1
	recursive = func(n int) int{
		if n == 0 {
			return memo[0]
		}
		if n == 1 {
			return memo[1]
		}
		if n == 2 {
			return memo[2]
		}
		if memo[n] == 0 {
			memo[n] = recursive(n-1) + recursive(n-2) + recursive(n-3) 
		}
		return memo[n]
	}

	return recursive(n)
}

/*
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. 
Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor.

Example 1:

Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.
Example 2:

Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.
 
Constraints:

2 <= cost.length <= 1000
0 <= cost[i] <= 999
*/
func minCostClimbingStairs(cost []int) int {
    var recursive func (arr []int, index int) int
	cache:=make([]int, len(cost))
	for i := range cache {
		cache[i] = -1
	}
	recursive = func(arr []int, index int) int {
		if index < 0 {
			return 0
		}
		if index == 0 {
			return arr[0]
		}
		if index == 1 {
			return arr[1]
		}
		
		if cache[index] == -1 {
			cache[index] = arr[index] + min(recursive(arr, index-1), recursive(arr, index-2))
		}
		return cache[index]
	}
	rs:=min(recursive(cost, len(cost)-1), recursive(cost,len(cost)-2))
	return rs
}

/*

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, 
the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
 

Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 400
*/
func rob(nums []int) int {
    var rec func (nums []int, index int) int
	cache:=make([]int, len(nums))
	for i:=0 ; i < len(nums);i++{
		cache[i] = -1
	}
	rec = func(nums []int, index int) int {
		if index < 0 {
			return 0
		}
		if index == 0 || index == 1 {
			return nums[index]
		}

		if cache[index] == -1 {
			cache[index] = max(rec(nums, index - 2), rec(nums, index-3)) + nums[index]
		}
		return cache[index]

	}

	rs:= max(rec(nums, len(nums)-1), rec(nums,len(nums)-2))
	return rs
}

/*
You have two types of tiles: a 2 x 1 domino shape and a tromino shape. You may rotate these shapes.
Given an integer n, return the number of ways to tile an 2 x n board. Since the answer may be very large, return it modulo 109 + 7.
In a tiling, every square must be covered by a tile. Two tilings are different if and only if there are two 4-directionally adjacent cells on the board such that exactly one of the tilings has both squares occupied by a tile.

Example 1:


Input: n = 3
Output: 5
Explanation: The five different ways are show above.
Example 2:

Input: n = 1
Output: 1
 
Constraints:

1 <= n <= 1000
*/
func numTilings(n int) int {
	dp:=make([]int, n +1)
	
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n== 3 {
		return 5
	}
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5
	mod:= 1_000_000_000 + 7


	for i:=4 ; i <= n; i ++ {
		dp[i] = 2*dp[i-1] + dp[i-3]
		dp[i] = dp[i] % mod
	}
	return dp[n]
}

/*
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). 
The robot can only move either down or right at any point in time.
Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.
The test cases are generated so that the answer will be less than or equal to 2 * 109.

Example 1:

Input: m = 3, n = 7
Output: 28
Example 2:

Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
 
Constraints:

1 <= m, n <= 100
*/
func uniquePaths(m int, n int) int {

	dp:= make([][]int, 0)
	for i:=0; i<m;i++ {
		tmp:= make([]int, n)
		for j:=0; j <n; j++ {
			tmp = append(tmp, 0)
		}
		dp = append(dp, tmp)
	}
	for i:=0 ; i < n; i++ {
		dp[0][i] = 1
	}
	for i:=0; i < m; i++ {
		dp[i][0] = 1
	}
	for i:=1 ; i < m; i++ {
		for j:=1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]

}

/*
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.
A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.

Example 1:

Input: text1 = "abcde", text2 = "ace" 
Output: 3  
Explanation: The longest common subsequence is "ace" and its length is 3.
Example 2:

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
Example 3:

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.
 

Constraints:

1 <= text1.length, text2.length <= 1000
text1 and text2 consist of only lowercase English characters.
*/
func longestCommonSubsequence(text1 string, text2 string) int {
    var recursive func(idx1, idx2 int) int
	fmt.Println()
	dp:=make([][]int,0)
	for i:=0 ; i <= len(text1); i++ {
		tmp:= make([]int, 0)
		for j:=0 ; j <= len(text2); j++ {
			tmp = append(tmp, -1)
		}
		dp = append(dp, tmp)
	}

	
	recursive = func(idx1, idx2 int) int {
		if idx1 == 0 || idx2 == 0 {
			return 0
		}

		if dp[idx1][idx2] != -1 {
			return dp[idx1][idx2]
		}
		if text1[idx1-1] == text2[idx2-1] {
			dp[idx1][idx2] = recursive(idx1-1, idx2-1) + 1
		}else {
			dp[idx1][idx2] = max(recursive(idx1-1, idx2), recursive(idx1, idx2-1))
		}
		return dp[idx1][idx2]
	}

	rs:= recursive(len(text1), len(text2))
	return rs
}

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day, and an integer fee representing a transaction fee.
Find the maximum profit you can achieve. 
You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.

Note:

You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
The transaction fee is only charged once for each stock purchase and sale.

Example 1:

Input: prices = [1,3,2,8,4,9], fee = 2
Output: 8
Explanation: The maximum profit can be achieved by:
- Buying at prices[0] = 1
- Selling at prices[3] = 8
- Buying at prices[4] = 4
- Selling at prices[5] = 9
The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
Example 2:

Input: prices = [1,3,7,5,10,3], fee = 3
Output: 6
 

Constraints:

1 <= prices.length <= 5 * 104
1 <= prices[i] < 5 * 104
0 <= fee < 5 * 104
*/
func maxProfit(prices []int, fee int) int {
    var recursive func(index int, isHold int) int
	dp := make([][]int,0)
	for i:=0 ; i< len(prices); i++ {
		tmp := make([]int,0)
		for j:=0 ; j < 2; j++ {
			tmp = append(tmp, -1)
		}
		dp = append(dp, tmp )
	}
	fmt.Println(dp)
	recursive = func(index int, isHold int) int {
		if index == len(prices) {
			return 0
		}

		if dp[index][isHold] != -1 {
			return dp[index][isHold]
		}

		if isHold == 1 {
			noSell:= recursive(index+1, 1)
			sell:= recursive(index+1, 0) + prices[index] - fee
			dp[index][isHold] = max(noSell, sell)
			return dp[index][isHold] 
		}else {
			buy:= recursive(index+1, 1) - prices[index] 
			noBuy:= recursive(index+1, 0) 
			dp[index][isHold] = max(buy, noBuy)
			return dp[index][isHold]
		}
	}

	rs:= recursive(0, 0)
	return rs
}

/*
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
You have the following three operations permitted on a word:
Insert a character
Delete a character
Replace a character

Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation: 
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation: 
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
 
Constraints:

0 <= word1.length, word2.length <= 500
word1 and word2 consist of lowercase English letters.
*/
func minDistance(word1 string, word2 string) int {
	if word1 =="" {
		return len(word2)
	}
	if word2 == "" {
		return len(word1)
	}
	var recursive func(index1, index2 int) int

	dp:= make([][]int, 0)
	for i:=0 ; i < len(word1); i++ {
		tmp:=make([]int,0)
		for j:=0; j < len(word2); j++ {
			tmp = append(tmp, -1)
		}
		dp = append(dp, tmp)
	}
	
	recursive = func(index1, index2 int) int {
		if index1 == len(word1) {
			return len(word2) - index2
		}
		if index2 == len(word2) {
			return len(word1) - index1
		}

		if dp[index1][index2] != -1 {
			return dp[index1][index2]
		}

		if word1[index1] == word2[index2] {
			dp[index1][index2] = recursive(index1+1, index2+1)
		}else  {
			replace := recursive(index1+1, index2+1)
			insert := recursive(index1, index2+1) 
			delete := recursive(index1+1, index2) 
			dp[index1][index2] = min(min(insert, delete), replace) + 1
		}
		return dp[index1][index2]
		
	}

	rs := recursive(0,0)
	return rs
}

/*
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), 
ans[i] is the number of 1's in the binary representation of i.

Example 1:

Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10
Example 2:

Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

Constraints:

0 <= n <= 10^5

Follow up:

It is very easy to come up with a solution with a runtime of O(n log n). Can you do it in linear time O(n) and possibly in a single pass?
Can you do it without using any built-in function (i.e., like __builtin_popcount in C++)?

// f(0)= 0000  
// f(1)= 0001   1

// f(2)= 0010   1
// f(3)= 0011   2

// f(4)= 0100  1

// f(5)= 0101  2
// f(6)= 0110  2
// f(7)= 0111  3

// f(8)= 1000 1

// f(9)= 1001 2
// f(10)= 1010 2
// f(11)= 1011 3
// f(12)= 1100 2
// f(13)= 1101 3
// f(14)= 1110 3
// f(15)= 1111 4

15/2 = 7


10/2 = 5 
10 - 5 = 5

// f(16)= 10000 1


3 = 2 + 1

1 + 3 = 4 

*/
func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	
	dp:=make([]int, n+1)
	if n== 1 {
		dp[1]=1
		return dp
	}

	dp[0]=0
	dp[1]=1
	binary:=1
	for i:=2; i<= n;i++ {
		if i < binary * 2 {
			dp[i]=dp[i-binary] + dp[binary]
		}else if i== binary*2{
			dp[i] = 1
		}else {
			binary = binary*2
			dp[i] = dp[i-binary] + dp[binary]
		}
	}

    // var recursive func(count int, binary int) int 
	// recursive = func(count int, binary int) int {
	// 	if binary == 0 {
	// 		return 0
	// 	}
	// 	if count == 1 || count == 2 {
	// 		return 1
	// 	}
	// 	if count == binary {
	// 		return 1
	// 	}

	// 	if count > binary {
	// 		return 1 + recursive(count - binary, binary)
	// 	}else {
	// 		return recursive(count, binary/2)
	// 	}
	// }

	// rs := recursive(n,i/2)
	// fmt.Println(rs)
	return dp
}