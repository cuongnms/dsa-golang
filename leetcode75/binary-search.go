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
	low:=1
	high:=n
    for {
		mid:=(high-low)/2 + low
		check := guess(mid)
		if check == 0 {
			return mid
		} else if check == -1 {
			high = mid
		}else {
			low = mid +1
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
	rs:=make([]int, 0)
    sort.Ints(potions)
	for i:=0; i < len(spells); i++ {
		tmp:=make([]int, len(potions))

		for j :=0 ; j < len(tmp); j++ {
			tmp[j] = potions[j] * spells[i]
		}
		rs = append(rs, getTotal(tmp, success))

	}
	return rs
}

func getTotal(arr []int, val int64) int {
	low:=0
	high := len(arr) - 1
	if int64(arr[low]) >= val {
		return len(arr)
	}
	if int64(arr[high]) < val {
		return 0
	}
	for {
		mid:=(high - low)/2 + low
		if low == mid-1 && int64(arr[mid]) >= val && int64(arr[low]) <= val {
			return len(arr) - mid
		}
		if mid + 1 == high && int64(arr[mid]) <= val && int64(arr[high]) >= val {
			return len(arr) - high
		}
		if int64(arr[mid]) > val {
			high = mid 
		}else if int64(arr[mid]) < val {
			low = mid
		} 
	}
}