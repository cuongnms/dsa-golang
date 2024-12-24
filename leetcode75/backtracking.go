package main

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Example 2:

Input: digits = ""
Output: []
Example 3:

Input: digits = "2"
Output: ["a","b","c"]

Constraints:

0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].
*/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	rs := make([]string, 0)

	var backtracking func(digits, tmp string, index int)

	backtracking = func(digits, tmp string, index int) {
		if index == len(digits) {
			rs = append(rs, tmp)
			return
		}
		checkArr := mapChar(digits[index])
		for i := 0; i < len(checkArr); i++ {
			tmp += checkArr[i]
			backtracking(digits, tmp, index+1)
			tmp = tmp[0 : len(tmp)-1]
		}
	}

	backtracking(digits, "", 0)
	return rs

}

func mapChar(char byte) []string {
	switch char {
	case '2':
		return []string{"a", "b", "c"}
	case '3':
		return []string{"d", "e", "f"}
	case '4':
		return []string{"g", "h", "i"}
	case '5':
		return []string{"j", "k", "l"}
	case '6':
		return []string{"m", "n", "o"}
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	default:
		return []string{}
	}
}

/*
Find all valid combinations of k numbers that sum up to n such that the following conditions are true:
Only numbers 1 through 9 are used.
Each number is used at most once.
Return a list of all possible valid combinations. The list must not contain the same combination twice,
and the combinations may be returned in any order.

Example 1:

Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.
Example 2:

Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Explanation:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.
Example 3:

Input: k = 4, n = 1
Output: []
Explanation: There are no valid combinations.
Using 4 different numbers in the range [1,9], the smallest sum we can get is 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.

Constraints:
2 <= k <= 9
1 <= n <= 60
*/
func combinationSum3(k int, n int) [][]int {
	rs := make([][]int, 0)
	tmp := make([]int, 0)
	// k = 3 , n = 7
	var backtracking func(comb []int, sum int, start int)
	backtracking = func(comb []int, sum int, start int) {
		if len(comb) > k {
			return
		}
		for i := start; i <= 9; i++ {
			sum += i
			comb = append(comb, i)
			if len(comb) == k && sum == n {
				clone := make([]int, len(comb))
				copy(clone, comb)
				rs = append(rs, clone)
				return
			}
			backtracking(comb, sum, i+1)
			comb = comb[0 : len(comb)-1]
			sum = sum - i
		}
	}
	backtracking(tmp, 0, 1)
	return rs
}
