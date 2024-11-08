package main

import "fmt"

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
    grid := make([][]int,2)
	set1:= NewSet()
	set2:= NewSet()
	for _, val:=range nums1 {
		if set1.Contains(val){
			continue
		}
		set1.Add(val)
	}
	for _, val:= range nums2 {
		if set2.Contains(val){
			continue
		}
		set2.Add(val)
	}

	for item:=range set1.items {
		if set2.Contains(item){
			set1.Remove(item)
			set2.Remove(item)
		}
	}
	rs1:= make([]int, 0)
	rs2:= make([]int, 0)
	for item:= range set1.items {
		rs1 = append(rs1, item)
	}
	for item:= range set2.items {
		rs2 = append(rs2, item)
	}
	grid[0] = rs1
	grid[1] = rs2
	return grid
}

