package main

import (
	"fmt"
	"strconv"
)

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
type ByteStack struct {
	data []byte
}

func NewByteStack() *ByteStack {
	return &ByteStack{}
}

func (s *ByteStack) Pop() (byte, error) {
	if len(s.data) == 0 {
		return byte(0), fmt.Errorf("empty stack")
	}
	rs := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return rs, nil
}

func (s *ByteStack) Push(letter byte) {
	s.data = append(s.data, letter)
}

func (s *ByteStack) Peek() (byte, error) {
	if len(s.data) == 0 {
		return byte(0), fmt.Errorf("empty stack")
	}
	return s.data[len(s.data)-1], nil
}

func (s *ByteStack) IsEmpty() bool {
	return len(s.data) == 0
}

func removeStars(s string) string {
	// leetcode
	//
	stack := NewByteStack()
	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			stack.Push(s[i])
		} else {
			stack.Pop()
		}
	}
	return string(stack.data)
}

/*
We are given an array asteroids of integers representing asteroids in a row.
For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left).
Each asteroid moves at the same speed.
Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode.
If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

Example 1:
[5, 10, -11, -12]

Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.
Example 2:

Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.
Example 3:

Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.

[-2, -1, 1, 2]

Constraints:

2 <= asteroids.length <= 104
-1000 <= asteroids[i] <= 1000
asteroids[i] != 0
*/
type IntStack struct {
	data []int
}

func NewIntStack() *IntStack {
	return &IntStack{}
}

func (s *IntStack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("empty stack")
	}
	rs := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return rs, nil
}

func (s *IntStack) Push(letter int) {
	s.data = append(s.data, letter)
}

func (s *IntStack) Peek() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("empty stack")
	}
	return s.data[len(s.data)-1], nil
}

func (s *IntStack) IsEmpty() bool {
	return len(s.data) == 0
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func asteroidCollision(asteroids []int) []int {
	stack := NewIntStack()
	i := 0
	for i < len(asteroids) {
		if stack.IsEmpty() {
			stack.Push(asteroids[i])
			i++
		} else {
			if val, err := stack.Peek(); err == nil && val*asteroids[i] < 0 {
				if val < 0 && asteroids[i] > 0 {
					stack.Push(asteroids[i])
					i++
					continue
				}
				if absInt(val) < absInt(asteroids[i]) {
					stack.Pop()
				} else if absInt(val) == absInt(asteroids[i]) {
					stack.Pop()
					i++
				} else {
					i++
				}
			} else if val, err := stack.Peek(); err == nil && val*asteroids[i] > 0 {

				stack.Push(asteroids[i])
				i++
			}
		}
	}
	return stack.data
}

/*
Given an encoded string, return its decoded string.
The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times.
Note that k is guaranteed to be a positive integer.
You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc.
Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k.
For example, there will not be input like 3a or 2[4].
The test cases are generated so that the length of the output will never exceed 105.

Example 1:
Input: s = "3[a]2[bc]"
Output: "aaabcbc" 

Example 2:
Input: s = "3[a2[c]]"
Output: "accaccacc"

// 3 [ a 2 [ c

Example 3:
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"

Constraints:
1 <= s.length <= 30
s consists of lowercase English letters, digits, and square brackets '[]'.
s is guaranteed to be a valid input.
All the integers in s are in the range [1, 300].
*/

type StrStack struct {
	data []string
}

func NewStrStack() *StrStack {
	return &StrStack{}
}

func (s *StrStack) Pop() (string, error) {
	if len(s.data) == 0 {
		return "", fmt.Errorf("empty stack")
	}
	rs := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return rs, nil
}

func (s *StrStack) Push(letter string) {
	s.data = append(s.data, letter)
}

func (s *StrStack) Peek() (string, error) {
	if len(s.data) == 0 {
		return "", fmt.Errorf("empty stack")
	}
	return s.data[len(s.data)-1], nil
}

func (s *StrStack) IsEmpty() bool {
	return len(s.data) == 0
}
func isDigit(b byte) bool {
    return b >= '0' && b <= '9'
}
func decodeString(s string) string {
	rs:=""
	countStack:= NewIntStack()
	rsStack := NewStrStack()
	index:=0
	for index < len(s) {
		if isDigit(s[index]) {
			strCount:=""
			for isDigit(s[index]) {
				strCount+=string(s[index])
				index++
			}
			if strCount != "" {
				count, _ := strconv.Atoi(strCount)
				countStack.Push(count)
			}
		}else if s[index] == '[' {
			rsStack.Push(rs)
			rs = ""
			index++
		} else if s[index] == ']' {
			tmpRs, _ := rsStack.Pop()
			repeatTime, _ := countStack.Pop()
			for i:=0; i < repeatTime; i++ {
				tmpRs+=rs
			}
			rs = tmpRs
			index++
		} else {
			rs+=string(s[index])
			index++
		}

	}
	return rs
}
