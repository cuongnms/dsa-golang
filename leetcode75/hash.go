package main

import "strconv"

// import "fmt"

/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.

Example 1:

Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Explanation:
For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].
Example 2:

Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
Output: [[3],[]]
Explanation:
For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
Every integer in nums2 is present in nums1. Therefore, answer[1] = [].

Constraints:

1 <= nums1.length, nums2.length <= 1000
-1000 <= nums1[i], nums2[i] <= 1000
*/

type Set struct {
	items map[int]struct{}
}

// NewSet initializes a new Set
func NewSet() *Set {
	return &Set{items: make(map[int]struct{})}
}

// Add inserts an element into the set
func (s *Set) Add(item int) {
	s.items[item] = struct{}{}
}

// Remove deletes an element from the set
func (s *Set) Remove(item int) {
	delete(s.items, item)
}

func (s *Set) Contains(item int) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set) Size() int {
	return len(s.items)
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	grid := make([][]int, 2)
	set1 := NewSet()
	set2 := NewSet()
	for _, val := range nums1 {
		if set1.Contains(val) {
			continue
		}
		set1.Add(val)
	}
	for _, val := range nums2 {
		if set2.Contains(val) {
			continue
		}
		set2.Add(val)
	}

	for item := range set1.items {
		if set2.Contains(item) {
			set1.Remove(item)
			set2.Remove(item)
		}
	}
	rs1 := make([]int, 0)
	rs2 := make([]int, 0)
	for item := range set1.items {
		rs1 = append(rs1, item)
	}
	for item := range set2.items {
		rs2 = append(rs2, item)
	}
	grid[0] = rs1
	grid[1] = rs2
	return grid
}

/*
Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.
Example 1:

Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
Example 2:

Input: arr = [1,2]
Output: false
Example 3:

Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true

Constraints:

1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000
*/

func uniqueOccurrences(arr []int) bool {

	mapOccurrences := make(map[int]int)
	for _, num := range arr {
		mapOccurrences[num]++
	}

	mapRs := make(map[int]int)

	for _, num := range mapOccurrences {
		if _, exist := mapRs[num]; exist {
			return false
		}
		mapRs[num]++
	}

	return true
}

/*
You can use the operations on either string as many times as necessary.

Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.

Example 1:

Input: word1 = "abc", word2 = "bca"
Output: true
Explanation: You can attain word2 from word1 in 2 operations.
Apply Operation 1: "abc" -> "acb"
Apply Operation 1: "acb" -> "bca"
Example 2:

Input: word1 = "a", word2 = "aa"
Output: false
Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.
Example 3:

Input: word1 = "cabbba", word2 = "abbccc"
Output: true
Explanation: You can attain word2 from word1 in 3 operations.
Apply Operation 1: "cabbba" -> "caabbb"
Apply Operation 2: "caabbb" -> "baaccc"
Apply Operation 2: "baaccc" -> "abbccc"

Constraints:

1 <= word1.length, word2.length <= 105
word1 and word2 contain only lowercase English letters.
*/
func closeStrings(word1 string, word2 string) bool {
	mapChar1 := make(map[byte]int)
	mapChar2 := make(map[byte]int)
	if len(word1) != len(word2) {
		return false
	}
	for i := 0; i < len(word1); i++ {
		mapChar1[word1[i]]++
		mapChar2[word2[i]]++
	}
	

	for key := range mapChar1 {
		if _, exist := mapChar2[key]; !exist {
			return false
		}
	}

	list1 := make(map[int]int)
	list2 := make(map[int]int)

	for _, val := range mapChar1 {
		list1[val]++
	}

	for _, val := range mapChar2 {
		list2[val]++
	}

	for key, val1 := range list1 {
		if val2, exist := list2[key]; !exist || val1 != val2 {
            return false
        }
	}

	return true
	
}


/*
Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

 

Example 1:


Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
Output: 1
Explanation: There is 1 equal row and column pair:
- (Row 2, Column 1): [2,7,7]
Example 2:


Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
Output: 3
Explanation: There are 3 equal row and column pairs:
- (Row 0, Column 0): [3,1,2,2]
- (Row 2, Column 2): [2,4,2,2]
- (Row 3, Column 2): [2,4,2,2]
 

Constraints:

n == grid.length == grid[i].length
1 <= n <= 200
1 <= grid[i][j] <= 105
*/
func equalPairs(grid [][]int) int {
    map1:=make(map[string]int)
	map2:=make(map[string]int)
	
	for r:=0; r<len(grid); r++{
		str:=""
		for c:=0 ; c<len(grid); c++ {
			str+="-" + strconv.Itoa(grid[r][c])
		}
		map1[str]++
	}

	for r:=0; r<len(grid); r++{
		str:=""
		for c:=0 ; c<len(grid); c++ {
			str+="-" + strconv.Itoa(grid[c][r])
		}
		map2[str]++
	}
	rs:=0
	for key1, val1:= range map1 {
		if val2, exist:=map2[key1]; exist {
            rs += val1 * val2
        }
	}
	return rs
}