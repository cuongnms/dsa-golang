package main


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
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	if len(nums) == k {
		return float64(sum) / float64(k)
	}
	max := sum
	for i := 1; i <= len(nums)-k; i++ {
		if max <= sum-nums[i-1]+nums[i+k-1] {
			max = sum - nums[i-1] + nums[i+k-1]
		}
		sum = sum - nums[i-1] + nums[i+k-1]
	}

	return float64(max) / float64(k)
}

/*
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

Example 1:

Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.
Example 2:

Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.
Example 3:

Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.

Constraints:

1 <= s.length <= 105
s consists of lowercase English letters.
1 <= k <= s.length
*/
func maxVowels(s string, k int) int {
	totalVowels := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			totalVowels++

		}
	}
	max := totalVowels
	for i := 1; i < len(s)-k+1; i++ {
		if isVowel(s[i+k-1]) {
			totalVowels++
		}
		if isVowel(s[i-1]) {
			totalVowels--
		}
		if totalVowels > max {
			max = totalVowels
		}
	}

	return max
}

func isVowel(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'u' || c == 'o' {
		return true
	}
	return false
}

/*
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

Example 1:

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]

Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
Example 2:

Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.
0 <= k <= nums.length
*/
func longestOnes(nums []int, k int) int {
	// left, right:=0,0
	// zeroCount:=0
	// max:= 0 
	// for right<len(nums) {
	// 	if nums[right] == 0 {
	// 		zeroCount++
	// 	}
	// 	for zeroCount > k {
	// 		if nums[left] == 0 {
	// 			zeroCount--
	// 		}
	// 		left++
	// 	}
	// 	if zeroCount <=k {
	// 		if max < right - left + 1 {
    //             max = right - left + 1
    //         }
	// 	}
	// 	right++
	// }
	// fmt.Println(max)
	// return max

	left, right:=0,0
	zeroCount:=0
	max:= 0 
	for right<len(nums) {
		if nums[right] == 0 {
			zeroCount++
		}
		if zeroCount > k {

			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		if zeroCount <= k {
			if max < right - left + 1 {
				max = right - left + 1
			}
		}
		right++
	}
    return max

}

/*
Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. 
Return 0 if there is no such subarray.

Example 1:

Input: nums = [1,1,0,1]
Output: 3
Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.
Example 2:

Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
Example 3:

Input: nums = [1,1,1]
Output: 2
Explanation: You must delete one element.
 
Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.
*/
func longestSubarray(nums []int) int {
	/*
	for right < len(nums) {
	  	if nums[right] == 1 => right++ => continue
		
		if nums[right] == 0 
			=> if zeroCount <= 1 
					=> zeroCount ++
					=> right++ 
					
					
	}
	*/
	left, right:=0,0
	z:=0
	max:=0
	for right < len(nums) {
		if nums[right] == 0 {
			z++
		}
		if z >1 {
			if nums[left] == 0 {
				z--
			}
			left++
		}
		if z<=1 {
			if max < right - left + 1 {
                max = right - left + 1
            }
		}
		right++
	}
	return max -1
}


