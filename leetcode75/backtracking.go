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
