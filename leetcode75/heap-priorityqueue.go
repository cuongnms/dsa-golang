package main

import (
	"container/heap"
	"fmt"
	"slices"
	"sort"
)

/*
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?



Example 1:

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Example 2:

Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4


Constraints:

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/

type Heap struct {
	arr  []int
}

func NewHeap() *Heap {
	arr := make([]int, 0)
	return &Heap{arr: arr}
}

func (h *Heap) heapifyUp(index int) {
	for index > 0 {
		parent:= (index-1)/2
		if h.arr[index]>=h.arr[parent]{
			break
		}
		h.arr[index], h.arr[parent] = h.arr[parent], h.arr[index]
		index = parent
	}
}

func (h *Heap) heapifyDown(index int) {
	lastIndex:=len(h.arr) - 1

	for {
		leftChild:=2*index+1
		rightChild:=2*index+2
		smallest:=index
		if leftChild <= lastIndex && h.arr[leftChild] < h.arr[smallest] {
			smallest = leftChild
		}

		if rightChild <= lastIndex && h.arr[rightChild] < h.arr[smallest] {
			smallest = rightChild
		}
		if smallest == index {
			break
		}
		h.arr[index], h.arr[smallest] = h.arr[smallest] , h.arr[index]
		index = smallest
	}

}

func (h *Heap) insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr)-1)
}

func (h *Heap) pop() int {
	if len(h.arr) ==0 {
		return -1
	}

	min:=h.arr[0]
	h.arr[0] = h.arr[len(h.arr) - 1]
	h.arr= h.arr[:len(h.arr)-1]
	h.heapifyDown(0)
	return min
}

func findKthLargest(nums []int, k int) int {
	b:=NewHeap()
	for i:=0; i <k;i++ {
		b.insert(nums[i])
	}
	for i:=k; i < len(nums);i++{
		if nums[i] > b.arr[0] {
			b.pop()
			b.insert(nums[i])
		}
	}
	return b.arr[0]
}

/*
You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].

Implement the SmallestInfiniteSet class:

SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain all positive integers.
int popSmallest() Removes and returns the smallest integer contained in the infinite set.
void addBack(int num) Adds a positive integer num back into the infinite set, if it is not already in the infinite set.
 

Example 1:

Input
["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
[[], [2], [], [], [], [1], [], [], []]
Output
[null, null, 1, 2, 3, null, 1, 4, 5]

Explanation
SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
smallestInfiniteSet.addBack(2);    // 2 is already in the set, so no change is made.
smallestInfiniteSet.popSmallest(); // return 1, since 1 is the smallest number, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 2, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 3, and remove it from the set.
smallestInfiniteSet.addBack(1);    // 1 is added back to the set.
smallestInfiniteSet.popSmallest(); // return 1, since 1 was added back to the set and
                                   // is the smallest number, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 4, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 5, and remove it from the set.
 
Constraints:

1 <= num <= 1000
At most 1000 calls will be made in total to popSmallest and addBack.
*/

type SmallestInfiniteSet struct {
    arr []int
}


func SmallestInfiniteSetConstructor() SmallestInfiniteSet {
	arr:= make([]int, 0)
	for i :=0; i<1000; i++{
		arr = append(arr, i+1)
	}
    return SmallestInfiniteSet{
		arr: arr,
	}
}

func (this *SmallestInfiniteSet) heapifyDownSet(index int) {
	lastIndex:= len(this.arr) - 1
	for {
		leftIndex:=2*index + 1
		rightIndex:=2*index + 2
		smallestValIndex:=index
		if leftIndex <= lastIndex && this.arr[leftIndex] < this.arr[smallestValIndex] {
            smallestValIndex = leftIndex
        }
		if rightIndex <= lastIndex && this.arr[rightIndex] < this.arr[smallestValIndex] {
            smallestValIndex = rightIndex
        }
		if smallestValIndex == index {
            break
        }
		this.arr[index], this.arr[smallestValIndex] = this.arr[smallestValIndex], this.arr[index]
		index= smallestValIndex

	}
}

func (this *SmallestInfiniteSet) heapifyUpSet(index int) {
	for index >0 {
		parentIndex:= (index-1)/2
		if this.arr[index] >= this.arr[parentIndex] {
			break
		}
		this.arr[index], this.arr[parentIndex] = this.arr[parentIndex], this.arr[index]
		index = parentIndex
	}
}


func (this *SmallestInfiniteSet) PopSmallest() int {
	rs:=this.arr[0]
	this.arr[0] = this.arr[len(this.arr)-1]
	this.arr= this.arr[:len(this.arr)-1]
	this.heapifyDownSet(0)
	return rs
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    if slices.Contains(this.arr, num) {
		return
	}else {
		this.arr = append(this.arr, num)
        this.heapifyUpSet(len(this.arr)-1)
	}
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

 /*
You are given two 0-indexed integer arrays nums1 and nums2 of equal length n and a positive integer k. 
You must choose a subsequence of indices from nums1 of length k.
For chosen indices i0, i1, ..., ik - 1, your score is defined as:
The sum of the selected elements from nums1 multiplied with the minimum of the selected elements from nums2.
It can defined simply as: (nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] , nums2[i1], ... ,nums2[ik - 1]).
Return the maximum possible score.

A subsequence of indices of an array is a set that can be derived from the set {0, 1, ..., n-1} by deleting some or no elements.

Example 1:

Input: nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3
Output: 12
Explanation: 
The four possible subsequence scores are:
- We choose the indices 0, 1, and 2 with score = (1+3+3) * min(2,1,3) = 7.
- We choose the indices 0, 1, and 3 with score = (1+3+2) * min(2,1,4) = 6. 
- We choose the indices 0, 2, and 3 with score = (1+3+2) * min(2,3,4) = 12. 
- We choose the indices 1, 2, and 3 with score = (3+3+2) * min(1,3,4) = 8.
Therefore, we return the max score, which is 12.
Example 2:

Input: nums1 = [4,2,3,1,1], nums2 = [7,5,10,9,6], k = 1
Output: 30
Explanation: 
Choosing index 2 is optimal: nums1[2] * nums2[2] = 3 * 10 = 30 is the maximum possible score.
 
Constraints:
n == nums1.length == nums2.length
1 <= n <= 105
0 <= nums1[i], nums2[j] <= 105
1 <= k <= n
 */

type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old:=*pq
	n:=len(old)
	item:=old[n-1]
	*pq=old[0:n-1]
	return item
}

 func maxScore(nums1 []int, nums2 []int, k int) int64 {
    pq:=&PriorityQueue{}
	heap.Init(pq)
	arr:=make([][2]int, 0)
	for i:=0; i < len(nums1);i++ {
		arr = append(arr, [2]int{nums1[i], nums2[i]})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	sum:=0
	var rs int64
	rs= 0 
	for i:=0; i<len(nums1); i++ {
		heap.Push(pq, arr[i][0])
		sum+= arr[i][0]
		if pq.Len()>k {
			val:= heap.Pop(pq).(int)
			sum-=val
		}
		if pq.Len() == k {
			fmt.Println(sum* arr[i][1])
			rs = maxInt64(rs, int64(sum* arr[i][1]))
		}
	}
	return rs
 }

 