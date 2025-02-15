package main

import "fmt"

/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104

Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/
func maxSubArray(nums []int) int {
    currentMax:=nums[0]
	max:=nums[0]
	startIndex:=0
	endIndex:=0
	for i:=1; i < len(nums); i++ {
		if currentMax + nums[i] > nums[i] {
			currentMax += nums[i]
		}else {
			currentMax = nums[i]
			startIndex = i
		}
		if max < currentMax {
			max = currentMax
			endIndex = i
		}
	}
	fmt.Println(startIndex," ", endIndex)
	return max
}


/*
Given a circular integer array nums of length n, return the maximum possible sum of a non-empty subarray of nums.

A circular array means the end of the array connects to the beginning of the array. Formally, the next element of nums[i] is nums[(i + 1) % n] and the previous element of nums[i] is nums[(i - 1 + n) % n].

A subarray may only include each element of the fixed buffer nums at most once. Formally, for a subarray nums[i], nums[i + 1], ..., nums[j], there does not exist i <= k1, k2 <= j with k1 % n == k2 % n.

Example 1:

Input: nums = [1,-2,3,-2]
Output: 3
Explanation: Subarray [3] has maximum sum 3.
Example 2:

Input: nums = [5,-3,5]
Output: 10
Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10.
Example 3:

Input: nums = [-3,-2,-3]
Output: -2
Explanation: Subarray [-2] has maximum sum -2.
 
Constraints:

n == nums.length
1 <= n <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
*/
func maxSubarraySumCircular(nums []int) int {
	sumArr:=nums[0]
	minSum:=nums[0]
	maxSum:=nums[0]
	currentMin:=nums[0]
	currentMax:=nums[0]

	for i:=1; i < len(nums); i++ {
		sumArr += nums[i]
		if currentMax + nums[i] > nums[i] {
			currentMax += nums[i]
		}else {
			currentMax = nums[i]
		}
		if maxSum < currentMax {
			maxSum = currentMax
		}

		if currentMin + nums[i] < nums[i] {
            currentMin += nums[i]
        }else {
            currentMin = nums[i]
        }
		if minSum > currentMin {
            minSum = currentMin
        }
	}
	if maxSum <= 0 {
		return maxSum
	}else {
		if maxSum > sumArr - minSum {
			return maxSum
		}else {
			return sumArr - minSum
		}
	}
}

