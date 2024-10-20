package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

/*
You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1.
If a string is longer than the other, append the additional letters onto the end of the merged string.
Return the merged string.

Example 1:

Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r
Example 2:

Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b
word2:    p   q   r   s
merged: a p b q   r   s
Example 3:

Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q
merged: a p b q c   d

Constraints:

1 <= word1.length, word2.length <= 100
word1 and word2 consist of lowercase English letters.
*/
func mergeAlternately(word1 string, word2 string) string {
	idx := 0
	rs := ""
	for idx < len(word1) && idx < len(word2) {
		rs = rs + string(word1[idx]) + string(word2[idx])
		idx++
	}

	if len(word1) > len(word2) {
		rs = rs + word1[idx:]
	} else {
		rs = rs + word2[idx:]
	}
	return rs
}

/*
For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""

Constraints:

1 <= str1.length, str2.length <= 1000
str1 and str2 consist of English uppercase letters.
*/
func gcdOfStrings(str1 string, str2 string) string {
	var rs string
	if len(str1) > len(str2) {
		rs = str2
	} else {
		rs = str1
	}
	for len(rs) > 0 {
		if isCommonSubString(str1, rs) && isCommonSubString(str2, rs) {
			return rs
		} else {
			rs = rs[0 : len(rs)-1]
			fmt.Println("rs " + rs)
		}
	}
	return ""
}

func isCommonSubString(word1 string, common string) bool {
	orgCommon := common
	if common == word1 {
		return true
	}
	for len(common) <= len(word1) {
		if common != word1[0:len(common)] {
			return false
		} else if common == word1 {
			return true
		} else {
			common = common + orgCommon
		}
	}
	return false
}

/*
There are n kids with candies. You are given an integer array candies,
where each candies[i] represents the number of candies the ith kid has,
and an integer extraCandies, denoting the number of extra candies that you have.
Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies,
they will have the greatest number of candies among all the kids, or false otherwise.
Note that multiple kids can have the greatest number of candies.

Example 1:

Input: candies = [2,3,5,1,3], extraCandies = 3
Output: [true,true,true,false,true]
Explanation: If you give all extraCandies to:
- Kid 1, they will have 2 + 3 = 5 candies, which is the greatest among the kids.
- Kid 2, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
- Kid 3, they will have 5 + 3 = 8 candies, which is the greatest among the kids.
- Kid 4, they will have 1 + 3 = 4 candies, which is not the greatest among the kids.
- Kid 5, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
Example 2:

Input: candies = [4,2,1,1,2], extraCandies = 1
Output: [true,false,false,false,false]
Explanation: There is only 1 extra candy.
Kid 1 will always have the greatest number of candies, even if a different kid is given the extra candy.
Example 3:

Input: candies = [12,1,12], extraCandies = 10
Output: [true,false,true]

Constraints:

n == candies.length
2 <= n <= 100
1 <= candies[i] <= 100
1 <= extraCandies <= 50
*/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := slices.Max(candies)
	var rs []bool
	for i := range candies {
		if candies[i]+extraCandies >= max {
			rs = append(rs, true)
		} else {
			rs = append(rs, false)
		}
	}
	return rs
}

/*
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty,
and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.


Example 1:
Input: flowerbed = [1,0,0,0,1], n = 1
Output: true

Example 2:
Input: flowerbed = [1,0,0,0,1], n = 2
Output: false

1,0,0,0,0,0,0

Constraints:

1 <= flowerbed.length <= 2 * 104
flowerbed[i] is 0 or 1.
There are no two adjacent flowers in flowerbed.
0 <= n <= flowerbed.length
*/

func canPlaceFlowers(flowerbed []int, n int) bool {

	if n==0 {
		return true
	}
	
	for i:=0; i < len(flowerbed); i++ {
		if (i==0 && flowerbed[i]==0 && flowerbed[i+1]==0) || 
		(i==len(flowerbed)- 1 && flowerbed[i] == 0 && flowerbed[i-1]==0) || 
		(i>0 && i < len(flowerbed) && flowerbed[i-1]==0 && flowerbed[i+1]==0 && flowerbed[i] ==0) {
			flowerbed[i]=1
			n--
			if n==0 {
				return true
			}
		} 
			
	}
	return false

	// if n == 0 {
	// 	return true
	// }
	// if len(flowerbed) == 1 {
	// 	if n <= 1 {
	// 		if flowerbed[0] == 1 {
	// 			return false
	// 		} else {
	// 			return true
	// 		}
	// 	} else {
	// 		return false
	// 	}
	// }

	// start, end := 0, 1
	// count := 0
	// for start < len(flowerbed) && end < len(flowerbed) {
	// 	if flowerbed[end] == 1 && flowerbed[start] == 1 {
	// 		count = count + (end-start)/2 - 1
	// 		start = end
	// 		end++
	// 	} else if flowerbed[end] == 1 {
	// 		count = count + (end-start)/2
	// 		start = end
	// 		end++
	// 	} else {
	// 		end++
	// 	}
	// }
	// if flowerbed[start] == 1 && flowerbed[end-1] == 0 {
	// 	count = count + (end-1-start)/2
	// } else if flowerbed[start] == 0 && flowerbed[end-1] == 0 {
	// 	count = count + (end-1-start)/2 + 1
	// }
	// if count >= n {
	// 	return true
	// } else {
	// 	return false
	// }
}

/*
Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

Example 1:

Input: s = "IceCreAm"

Output: "AceCreIm"

Explanation:

The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".

Example 2:

Input: s = "leetcode"

Output: "leotcede"

Constraints:

1 <= s.length <= 3 * 105
s consist of printable ASCII characters.
*/
func reverseVowels(s string) string {
	if len(s) <= 1 {
		return s
	}
	rune := []rune(s)
	left, right := 0, len(rune)-1
	for left < right {
		if (unicode.ToLower(rune[left]) == 'a' || unicode.ToLower(rune[left]) == 'e' ||
			unicode.ToLower(rune[left]) == 'i' || unicode.ToLower(rune[left]) == 'u' ||
			unicode.ToLower(rune[left]) == 'o') && (unicode.ToLower(rune[right]) == 'a' || unicode.ToLower(rune[right]) == 'e' ||
			unicode.ToLower(rune[right]) == 'i' || unicode.ToLower(rune[right]) == 'u' ||
			unicode.ToLower(rune[right]) == 'o') {
			rune[left], rune[right] = rune[right], rune[left]
			left++
			right--
		} else if unicode.ToLower(rune[left]) != 'a' && unicode.ToLower(rune[left]) != 'e' && unicode.ToLower(rune[left]) != 'i' &&
			unicode.ToLower(rune[left]) != 'u' && unicode.ToLower(rune[left]) != 'o' {
			left++
		} else if unicode.ToLower(rune[right]) != 'a' && unicode.ToLower(rune[right]) != 'e' && unicode.ToLower(rune[right]) != 'i' &&
			unicode.ToLower(rune[right]) != 'u' && unicode.ToLower(rune[right]) != 'o' {
			right--
		}
	}

	return string(rune)
}

/*
Given an input string s, reverse the order of the words.
A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
Return a string of the words in reverse order concatenated by a single space.
Note that s may contain leading or trailing spaces or multiple spaces between two words.
The returned string should only have a single space separating the words. Do not include any extra spaces.

Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

Constraints:

1 <= s.length <= 104
s contains English letters (upper-case and lower-case), digits, and spaces ' '.
There is at least one word in s.

Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?
*/
func reverseWords(s string) string {
	f := func(c rune) bool {
		return c == ' '
	}
	arr := strings.FieldsFunc(s, f)

	left, right := 0, len(arr)-1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	return strings.Join(arr, " ")
}

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3,0]
Output: [0,0,9,0,0,0]

Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.


Follow up: Can you solve the problem in O(1) extra space complexity?
(The output array does not count as extra space for space complexity analysis.)
*/

func productExceptSelf(nums []int) []int {
	//Approach 1: calculate product then divine to each element in array
	// product := 1
	// hasZero:= false
	// countZero := 0
	// for _, value := range nums {
	// 	if value == 0 {
	// 		countZero++
	// 		hasZero = true
	// 	} else {
	// 		product = product * value
	// 	}
	// }
	// var rs []int
	// if countZero > 1 {
	// 	for i := 0; i < len(nums); i++ {
	// 		rs = append(rs, 0)
	// 	}
	// 	return rs
	// } else {
	// 	for i := 0; i < len(nums); i++ {
	// 		if nums[i] != 0 {
	// 			if hasZero {
	// 				rs = append(rs, 0)
	// 			}else {
	// 				rs = append(rs, product/nums[i])
	// 			}

	// 		} else {
	// 			rs = append(rs, product)
	// 		}
	// 	}
	// 	return rs
	// }
	//Approach 2: calculate the suffix and prefix of each element in array
	//Example:
	//            [1, 2 ,4, 6 ,8]
	//-> prefix:  [1, 1, 1*2, 1*2*4, 1*2*4*6]
	//-> suffix:  [2*4*6*8, 4*6*8, 6*8, 8, 1]
	var prefix []int
	var suffix []int
	for i := 0; i < len(nums); i++ {
		prefix = append(prefix, 1)
		suffix = append(suffix, 1)
	}

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	fmt.Println(prefix)
	fmt.Println(suffix)

	var rs []int
	for i := 0; i < len(nums); i++ {
		rs = append(rs, prefix[i]*suffix[i])
	}
	return rs
}

/*
Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k].
If no such indices exists, return false.

Example 1:

Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
Example 2:

Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
Example 3:

Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.

Constraints:

1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1
Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?
*/

func increasingTriplet(nums []int) bool {
	min1 := math.MaxInt
	min2 := math.MaxInt

	for _, num := range nums {
		if num <= min1 {
			min1 = num
		} else if num <= min2 {
			min2 = num

		} else {
			return true
		}
	}
	return false
}

/*
Given an array of characters chars, compress it using the following algorithm:

Begin with an empty string s. For each group of consecutive repeating characters in chars:

If the group's length is 1, append the character to s.
Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead, be stored in the input character array chars.
Note that group lengths that are 10 or longer will be split into multiple characters in chars.

After you are done modifying the input array, return the new length of the array.

You must write an algorithm that uses only constant extra space.

Example 1:

Input: chars = ["a","a","b","b","c","c","c"]
Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".
Example 2:

Input: chars = ["a"]
Output: Return 1, and the first character of the input array should be: ["a"]
Explanation: The only group is "a", which remains uncompressed since it's a single character.
Example 3:				
			     a   b  
Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".

Constraints:

1 <= chars.length <= 2000
chars[i] is a lowercase English letter, uppercase English letter, digit, or symbol.
*/
func compress(chars []byte) int {

	// a a b b b b c a a
    // a 2 b 4 c a 2 a a
	// a2b4ca2
	if len(chars) == 1 {
		return 1
	}
	pointer := 0
	i:=0
	iRs := 0
	for i < len(chars) {
		for pointer < len(chars) {
			if chars[i] == chars[pointer] {
				pointer++
			}else {
				break
			}
		}
		if pointer - i > 1 {
			rsStr := string(chars[i]) + strconv.Itoa(pointer-i)
			for j:=0; j < len(rsStr); j++ {
				chars[iRs+j] = rsStr[j]
			}
			iRs=iRs+len(rsStr)
		}else {
			chars[iRs] = chars[i]
			iRs=iRs+1
		}
		i = pointer
	}

	return iRs
}

