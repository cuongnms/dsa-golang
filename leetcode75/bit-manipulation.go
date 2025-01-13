package main

import (
	"fmt"
)

/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.



Example 1:

Input: nums = [2,2,1]

Output: 1

Example 2:

Input: nums = [4,1,2,1,2]

Output: 4

Example 3:

Input: nums = [1]

Output: 1

Constraints:

1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
Each element in the array appears twice except for one element which appears only once.
*/
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	tmp:=0
	for i:=0 ; i < len(nums); i++ {
		tmp = tmp^nums[i]
	}
	return tmp
}

/*
Given 3 positives numbers a, b and c. Return the minimum flips required in some bits of a and b to make ( a OR b == c ). (bitwise OR operation).
Flip operation consists of change any single bit 1 to 0 or change the bit 0 to 1 in their binary representation.

Example 1:

Input: a = 2, b = 6, c = 5
Output: 3
Explanation: After flips a = 1 , b = 4 , c = 5 such that (a OR b == c)
Example 2:

Input: a = 4, b = 2, c = 7
Output: 1
Example 3:

Input: a = 1, b = 2, c = 3
Output: 0
 
Constraints:

1 <= a <= 10^9
1 <= b <= 10^9
1 <= c <= 10^9


01000101

00001000
=
01001101



00010110
*/
func minFlips(a int, b int, c int) int {
    binarya := fmt.Sprintf("%0*b", 30, a)
	binaryb := fmt.Sprintf("%0*b", 30, b)
	binaryc := fmt.Sprintf("%0*b", 30, c)
	count:=0
	for i:=0 ; i < len(binaryc); i++ {
		if binaryc[i] == '0' {
			if binarya[i] != '0'{
				count++
			}

			if binaryb[i] != '0'{
				count++
			}
		}else {
			if binarya[i] == '0' && binaryb[i] == '0'{
				count++
			}
		}

	}

	
	return count
}