package main

import (
	"fmt"
	"sort"
)

/*
We are playing the Guess Game. The game is as follows:
I pick a number from 1 to n. You have to guess which number I picked.
Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.
You call a pre-defined API int guess(int num), which returns three possible results:

-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.

Example 1:

Input: n = 10, pick = 6
Output: 6
Example 2:

Input: n = 1, pick = 1
Output: 1
Example 3:

Input: n = 2, pick = 1
Output: 1

Constraints:

1 <= n <= 231 - 1
1 <= pick <= n
*/

func guess(number int) int {
	fmt.Println()
	return 1
}
func guessNumber(n int) int {
	low := 1
	high := n
	for {
		mid := (high-low)/2 + low
		check := guess(mid)
		if check == 0 {
			return mid
		} else if check == -1 {
			high = mid
		} else {
			low = mid + 1
		}
	}
}

/*
You are given two positive integer arrays spells and potions, of length n and m respectively,
where spells[i] represents the strength of the ith spell and potions[j] represents the strength of the jth potion.

You are also given an integer success. A spell and potion pair is considered successful if the product of their strengths is at least success.

Return an integer array pairs of length n where pairs[i] is the number of potions that will form a successful pair with the ith spell.

Example 1:

Input: spells = [5,1,3], potions = [1,2,3,4,5], success = 7
Output: [4,0,3]
Explanation:
- 0th spell: 5 * [1,2,3,4,5] = [5,10,15,20,25]. 4 pairs are successful.
- 1st spell: 1 * [1,2,3,4,5] = [1,2,3,4,5]. 0 pairs are successful.
- 2nd spell: 3 * [1,2,3,4,5] = [3,6,9,12,15]. 3 pairs are successful.
Thus, [4,0,3] is returned.
Example 2:

Input: spells = [3,1,2], potions = [8,5,8], success = 16
Output: [2,0,2]
Explanation:
- 0th spell: 3 * [8,5,8] = [24,15,24]. 2 pairs are successful.
- 1st spell: 1 * [8,5,8] = [8,5,8]. 0 pairs are successful.
- 2nd spell: 2 * [8,5,8] = [16,10,16]. 2 pairs are successful.
Thus, [2,0,2] is returned.

Constraints:

n == spells.length
m == potions.length
1 <= n, m <= 105
1 <= spells[i], potions[i] <= 105
1 <= success <= 1010
*/
func successfulPairs(spells []int, potions []int, success int64) []int {
	rs := make([]int, 0)
	sort.Ints(potions)
	for i := 0; i < len(spells); i++ {
		low := 0
		high := len(potions) - 1

		for low <= high {
			mid := low + (high-low)/2
			product := spells[i] * potions[mid]
			if product >= int(success) {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		rs = append(rs, len(potions)-low)

	}
	return rs
}

/*
A peak element is an element that is strictly greater than its neighbors.

Given a 0-indexed integer array nums, find a peak element, and return its index.
If the array contains multiple peaks, return the index to any of the peaks.
You may imagine that nums[-1] = nums[n] = -∞.
In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.
You must write an algorithm that runs in O(log n) time.

Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
Example 2:

Input: nums = [1,2,1,3,5,6,4]
Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.

Constraints:

1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
nums[i] != nums[i + 1] for all valid i.
*/
func findPeakElement(nums []int) int {
	rs := findRecursive(nums, 0, len(nums)-1)
	return rs
}

func findRecursive(nums []int, low, high int) int {
	if low == high {
		return high
	}

	if low+1 == high {
		if nums[low] >= nums[high] {
			return low
		} else {
			return high
		}
	}

	mid := (high-low)/2 + low
	max1 := findRecursive(nums, mid, high)
	max2 := findRecursive(nums, low, mid-1)
	if nums[max1] >= nums[max2] {
		return max1
	} else {
		return max2
	}
}

/*
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. 
If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.

Example 1:

Input: piles = [3,6,7,11], h = 8
Output: 4
Example 2:

Input: piles = [30,11,23,4,20], h = 5
Output: 30
Example 3:

Input: piles = [30,11,23,4,20], h = 6
Output: 23
 
Constraints:

1 <= piles.length <= 104
piles.length <= h <= 109
1 <= piles[i] <= 109
*/
func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	low := 1
	high := piles[len(piles) - 1]
	for low <= high {
		k:= (low + high) /2
		if checkTime(piles, k, h) {
			high = k - 1
		}else {
			low = k + 1
		}
	}
	return low
}

func checkTime(arr []int, k,h int ) bool {
	time:=0
	for i:=0 ; i < len(arr); i++ {
		val:= arr[i]/k
		if val == 0 {
			time++
		}else {
			bonus:= arr[i]%k
			time+= val
			if bonus > 0 {
				time++
			}
		}
		
		
	}
	return time <= h 
}