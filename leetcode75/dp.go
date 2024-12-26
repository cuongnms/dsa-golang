package main


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