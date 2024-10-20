package main

import "fmt"

/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.



Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
Example 2:

Input: nums = [5], k = 1
Output: 5.00000


Constraints:

n == nums.length
1 <= k <= n <= 105
-104 <= nums[i] <= 104
*/
func findMaxAverage(nums []int, k int) float64 {
	sum:=0
    for i:=0; i <k; i++ {
		sum+=nums[i]
	}
	if len(nums)==k {
		return float64(sum) / float64(k)
	}
	max:=sum
	for i:=1; i <= len(nums)-k; i++ {
		if max <= sum - nums[i-1] + nums[i+k-1] { 
			max = sum - nums[i-1] + nums[i+k-1]
		}
		sum = sum - nums[i-1] + nums[i+k-1]
	}

	return float64(max)/float64(k)
}