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