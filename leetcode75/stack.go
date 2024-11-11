package main

import "fmt"

/*
You are given a string s, which contains stars *.
In one operation, you can:
Choose a star in s.
Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.
Note:

The input will be generated such that the operation is always possible.
It can be shown that the resulting string will always be unique.

Example 1:
Input: s = "leet**cod*e"
Output: "lecoe"
Explanation: Performing the removals from left to right:
- The closest character to the 1st star is 't' in "leet**cod*e". s becomes "lee*cod*e".
- The closest character to the 2nd star is 'e' in "lee*cod*e". s becomes "lecod*e".
- The closest character to the 3rd star is 'd' in "lecod*e". s becomes "lecoe".
There are no more stars, so we return "lecoe".
Example 2:

Input: s = "erase*****"
Output: ""
Explanation: The entire string is removed, so we return an empty string.

Constraints:
1 <= s.length <= 105
s consists of lowercase English letters and stars *.
The operation above can be performed on s.
*/
type Stack struct{
	data []byte
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() (byte, error) {
	if len(s.data) == 0 {
		return byte(0), fmt.Errorf("Empty stack") 
	}
	rs:= s.data[len(s.data)- 1]
	s.data = s.data[:len(s.data)-1]
	return rs, nil
}

func (s *Stack) Push(letter byte) {
	s.data = append(s.data, letter)
}

func (s *Stack) Peek() (byte, error) {
	if len(s.data) == 0 {
		return byte(0), fmt.Errorf("Empty stack")
	}
	return s.data[len(s.data) - 1], nil
}

func removeStars(s string) string {
    // leetcode
	// 
	stack := NewStack()
	for i:=0; i < len(s); i++ {
		if s[i] != '*' {
			stack.Push(s[i])
		}else {
			stack.Pop()
		}
	}
	return string(stack.data)
}